// haai - Human Activity Automation Index CLI
// A tool for querying the HAAI taxonomy data
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Data structures for taxonomy
type Taxonomy struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	ID                    int        `json:"id"`
	Name                  string     `json:"name"`
	Description           string     `json:"description"`
	AbstractionScore      int        `json:"abstractionScore"`
	EstimatedAgiWave      any        `json:"estimatedAgiWave"` // can be int or []int
	PrimaryAISystemType   string     `json:"primaryAiSystemType"`
	Categories            []Category `json:"categories"`
}

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Index file structure (v2.0.0 data model)
type IndexFile struct {
	IndexID     string            `json:"indexId"`
	IndexName   string            `json:"indexName"`
	Description string            `json:"description"`
	Version     string            `json:"version"`
	Scale       IndexScale        `json:"scale"`
	Values      map[string]int    `json:"values"`
}

type IndexScale struct {
	Min    int          `json:"min"`
	Max    int          `json:"max"`
	Levels []IndexLevel `json:"levels"`
}

type IndexLevel struct {
	Level      int    `json:"level"`
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

// Activities file structure (v2.0.0)
type ActivitiesFile struct {
	Version         string     `json:"version"`
	ActivitiesCount int        `json:"activitiesCount"`
	Activities      []Activity `json:"activities"`
}

// Assessment file structure
type AssessmentFile struct {
	AssessmentDate      string                     `json:"assessmentDate"`
	Version             string                     `json:"version"`
	AGIWaveTimelines    AGIWaveTimelines           `json:"agiWaveTimelines"`
	ActivityAssessments map[string]json.RawMessage `json:"activityAssessments"`
}

type AGIWaveTimelines struct {
	Description string     `json:"description"`
	Waves       []WaveInfo `json:"waves"`
}

type WaveInfo struct {
	Wave            int      `json:"wave"`
	Timeline        string   `json:"timeline"`
	Characteristics string   `json:"characteristics"`
	TypicalDomains  []string `json:"typicalDomains"`
	KeyCapabilities []string `json:"keyCapabilities"`
}

type ActivityAssessment struct {
	AICapability string `json:"aiCapability"`
	Bottleneck   string `json:"bottleneck"`
	AGIWave      int    `json:"agiWave"`
}

// Activity from domain files
type DomainFile struct {
	DomainID   int        `json:"domainId"`
	DomainName string     `json:"domainName"`
	Activities []Activity `json:"activities"`
}

type Activity struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	DomainID     int      `json:"domainId"`
	CategoryID   string   `json:"categoryId"`
	Scores       Scores   `json:"scores"`
	ExampleTasks []string `json:"exampleTasks"`
}

type Scores struct {
	Abstraction      int    `json:"abstraction"`
	ErrorTolerance   int    `json:"errorTolerance"`
	FeedbackSpeed    int    `json:"feedbackSpeed"`
	SocialComplexity int    `json:"socialComplexity"`
	Purpose          int    `json:"purpose"`
	AICapability     string `json:"aiCapability"`
	Bottleneck       string `json:"bottleneck"`
	AGIWave          int    `json:"agiWave"`
}

// Mappings data
type Mappings struct {
	ATUSMapping    ATUSMapping    `json:"atusMapping"`
	EconomicImpact EconomicImpact `json:"economicImpact"`
}

type ATUSMapping struct {
	DataSource string        `json:"dataSource"`
	TimeUnit   string        `json:"timeUnit"`
	Mappings   []ATUSEntry   `json:"mappings"`
	Summary    ATUSSummary   `json:"summary"`
}

type ATUSEntry struct {
	ATUSCode          string             `json:"atusCode"`
	ATUSCategory      string             `json:"atusCategory"`
	HAICategories     []string           `json:"haaiCategories"`
	Notes             string             `json:"notes"`
	AvgMinutesPerDay  int                `json:"avgMinutesPerDay"`
	ParticipationRate float64            `json:"participationRate"`
	Breakdown         map[string]int     `json:"breakdown,omitempty"`
}

type ATUSSummary struct {
	TotalMinutesPerDay   int `json:"totalMinutesPerDay"`
	SleepAndPersonalCare int `json:"sleepAndPersonalCare"`
	Work                 int `json:"work"`
	Leisure              int `json:"leisure"`
	HouseholdAndCare     int `json:"householdAndCare"`
	Travel               int `json:"travel"`
	Other                int `json:"other"`
}

type EconomicImpact struct {
	Description   string           `json:"description"`
	Currency      string           `json:"currency"`
	Year          int              `json:"year"`
	USLaborMarket USLaborMarket    `json:"usLaborMarket"`
	DomainEcon    []DomainEconomic `json:"domainEconomics"`
}

type USLaborMarket struct {
	TotalEmployment   int     `json:"totalEmployment"`
	TotalWages        int64   `json:"totalWages"`
	AverageHourlyWage float64 `json:"averageHourlyWage"`
}

type DomainEconomic struct {
	DomainID            int      `json:"domainId"`
	DomainName          string   `json:"domainName"`
	EstimatedWorkers    int      `json:"estimatedWorkers"`
	PercentOfWorkforce  float64  `json:"percentOfWorkforce"`
	MedianHourlyWage    *float64 `json:"medianHourlyWage"`
	AnnualValueBillions int      `json:"annualValueBillions"`
	AutomationExposure  string   `json:"automationExposure"`
	SampleOccupations   []string `json:"sampleOccupations"`
	Notes               string   `json:"notes"`
}

var basePath string

func init() {
	// Find the base path (where taxonomy.json is)
	exe, err := os.Executable()
	if err == nil {
		// Try relative to executable
		dir := filepath.Dir(exe)
		if _, err := os.Stat(filepath.Join(dir, "..", "..", "taxonomy.json")); err == nil {
			basePath = filepath.Join(dir, "..", "..")
			return
		}
	}
	// Try current directory
	if _, err := os.Stat("taxonomy.json"); err == nil {
		basePath = "."
		return
	}
	// Try parent directories
	cwd, _ := os.Getwd()
	for dir := cwd; dir != "/"; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, "taxonomy.json")); err == nil {
			basePath = dir
			return
		}
	}
	basePath = "."
}

func loadJSON(filename string, v any) error {
	path := filepath.Join(basePath, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	return json.Unmarshal(data, v)
}

func loadTaxonomy() (*Taxonomy, error) {
	var t Taxonomy
	if err := loadJSON("taxonomy.json", &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func loadActivities() ([]Activity, error) {
	// First try to load from new activities.json format
	var af ActivitiesFile
	if err := loadJSON("activities.json", &af); err == nil && len(af.Activities) > 0 {
		// Successfully loaded v2.0.0 format
		activities := af.Activities

		// Load indices and assessments to populate scores
		indices, _ := loadIndices()
		assessment, _ := loadLatestAssessment()

		// For feedbackSpeed and socialComplexity, we need the old domain files
		feedbackMap, socialMap := loadLegacyScores()

		// Merge scores into activities
		mergeScores(activities, indices, assessment, feedbackMap, socialMap)

		return activities, nil
	}

	// Fall back to old domain-X.json files for backwards compatibility
	var all []Activity
	for i := 1; i <= 10; i++ {
		var df DomainFile
		filename := fmt.Sprintf("activities/domain-%d.json", i)
		if err := loadJSON(filename, &df); err != nil {
			continue // Skip missing files
		}
		// Set DomainID for each activity from the domain file
		for j := range df.Activities {
			df.Activities[j].DomainID = df.DomainID
		}
		all = append(all, df.Activities...)
	}
	return all, nil
}

// loadIndices loads abstraction.json, error-tolerance.json, and purpose.json from indices/
func loadIndices() (map[string]*IndexFile, error) {
	indices := make(map[string]*IndexFile)
	indexNames := []string{"abstraction", "error-tolerance", "purpose"}

	for _, name := range indexNames {
		var idx IndexFile
		filename := fmt.Sprintf("indices/%s.json", name)
		if err := loadJSON(filename, &idx); err != nil {
			continue // Skip missing files
		}
		indices[name] = &idx
	}

	return indices, nil
}

// loadLatestAssessment finds and loads the most recent assessment file from assessments/
func loadLatestAssessment() (*AssessmentFile, error) {
	assessmentsDir := filepath.Join(basePath, "assessments")
	entries, err := os.ReadDir(assessmentsDir)
	if err != nil {
		return nil, err
	}

	// Find the most recent assessment file (sorted by name = date)
	var latestFile string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			if entry.Name() > latestFile {
				latestFile = entry.Name()
			}
		}
	}

	if latestFile == "" {
		return nil, fmt.Errorf("no assessment files found")
	}

	var af AssessmentFile
	if err := loadJSON("assessments/"+latestFile, &af); err != nil {
		return nil, err
	}

	return &af, nil
}

// loadLegacyScores loads feedbackSpeed and socialComplexity from old domain files
func loadLegacyScores() (map[string]int, map[string]int) {
	feedbackMap := make(map[string]int)
	socialMap := make(map[string]int)

	for i := 1; i <= 10; i++ {
		var df DomainFile
		filename := fmt.Sprintf("activities/domain-%d.json", i)
		if err := loadJSON(filename, &df); err != nil {
			continue
		}
		for _, a := range df.Activities {
			feedbackMap[a.ID] = a.Scores.FeedbackSpeed
			socialMap[a.ID] = a.Scores.SocialComplexity
		}
	}

	return feedbackMap, socialMap
}

// mergeScores merges index values and assessment data into activity Scores structs
func mergeScores(activities []Activity, indices map[string]*IndexFile, assessment *AssessmentFile, feedbackMap, socialMap map[string]int) {
	for i := range activities {
		id := activities[i].ID

		// Merge abstraction index
		if idx, ok := indices["abstraction"]; ok {
			if val, ok := idx.Values[id]; ok {
				activities[i].Scores.Abstraction = val
			}
		}

		// Merge error-tolerance index
		if idx, ok := indices["error-tolerance"]; ok {
			if val, ok := idx.Values[id]; ok {
				activities[i].Scores.ErrorTolerance = val
			}
		}

		// Merge purpose index
		if idx, ok := indices["purpose"]; ok {
			if val, ok := idx.Values[id]; ok {
				activities[i].Scores.Purpose = val
			}
		}

		// Merge feedbackSpeed and socialComplexity from legacy scores
		if val, ok := feedbackMap[id]; ok {
			activities[i].Scores.FeedbackSpeed = val
		}
		if val, ok := socialMap[id]; ok {
			activities[i].Scores.SocialComplexity = val
		}

		// Merge assessment data
		if assessment != nil {
			if raw, ok := assessment.ActivityAssessments[id]; ok {
				var aa ActivityAssessment
				if err := json.Unmarshal(raw, &aa); err == nil {
					activities[i].Scores.AICapability = aa.AICapability
					activities[i].Scores.Bottleneck = aa.Bottleneck
					activities[i].Scores.AGIWave = aa.AGIWave
				}
			}
		}
	}
}

func loadMappings() (*Mappings, error) {
	var m Mappings
	if err := loadJSON("mappings.json", &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func printUsage() {
	fmt.Println(`haai - Human Activity Automation Index CLI

Usage:
  haai <command> [arguments]

Commands:
  domains              List all 10 domains with abstraction scores
  domain <id>          Show domain details and its categories
  activities [domain]  List activities (optionally filter by domain ID)
  activity <id>        Show activity details (e.g., "3.3.1")
  wave <n>             List activities by AGI wave (1-4)
  capability <status>  List by AI capability (solved, near_solved, partial, early, not_attempted)
  bottleneck <type>    List by bottleneck (dexterity, social, reasoning, mobility, etc.)
  index <name>         Show details about an index (abstraction, error-tolerance, or purpose)
  purpose <level>      List activities by purpose level (1-5)
  search <term>        Search activities by name or description
  time                 Show ATUS time-spent data
  econ                 Show economic impact by domain
  stats                Show summary statistics

Examples:
  haai domains
  haai domain 3
  haai activities 1
  haai activity 3.3.1
  haai wave 1
  haai capability solved
  haai bottleneck dexterity
  haai index purpose
  haai purpose 1
  haai search "code"
  haai time
  haai econ
  haai stats`)
}

func cmdDomains() {
	tax, err := loadTaxonomy()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("HAAI Domains (ordered by abstraction level)")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("%-4s %-35s %-6s %-8s\n", "ID", "Domain", "Abstr", "AGI Wave")
	fmt.Println(strings.Repeat("-", 70))

	for _, d := range tax.Domains {
		wave := formatWave(d.EstimatedAgiWave)
		fmt.Printf("%-4d %-35s %-6d %-8s\n", d.ID, d.Name, d.AbstractionScore, wave)
	}
}

func formatWave(w any) string {
	switch v := w.(type) {
	case float64:
		return fmt.Sprintf("%d", int(v))
	case int:
		return fmt.Sprintf("%d", v)
	case []any:
		parts := make([]string, len(v))
		for i, x := range v {
			parts[i] = fmt.Sprintf("%v", x)
		}
		return strings.Join(parts, "-")
	default:
		return fmt.Sprintf("%v", w)
	}
}

func cmdDomain(id int) {
	tax, err := loadTaxonomy()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var domain *Domain
	for _, d := range tax.Domains {
		if d.ID == id {
			domain = &d
			break
		}
	}
	if domain == nil {
		fmt.Fprintf(os.Stderr, "Domain %d not found\n", id)
		os.Exit(1)
	}

	fmt.Printf("Domain %d: %s\n", domain.ID, domain.Name)
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Description:     %s\n", domain.Description)
	fmt.Printf("Abstraction:     %d/10\n", domain.AbstractionScore)
	fmt.Printf("AGI Wave:        %s\n", formatWave(domain.EstimatedAgiWave))
	fmt.Printf("Primary AI:      %s\n", domain.PrimaryAISystemType)
	fmt.Println()
	fmt.Println("Categories:")
	for _, c := range domain.Categories {
		fmt.Printf("  %s: %s\n", c.ID, c.Name)
		fmt.Printf("      %s\n", c.Description)
	}
}

func cmdActivities(domainFilter int) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if domainFilter > 0 {
		fmt.Printf("Activities in Domain %d\n", domainFilter)
	} else {
		fmt.Println("All Activities")
	}
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-12s %-5s\n", "ID", "Name", "Capability", "Wave")
	fmt.Println(strings.Repeat("-", 80))

	for _, a := range activities {
		domainID := getDomainFromID(a.ID)
		if domainFilter > 0 && domainID != domainFilter {
			continue
		}
		name := a.Name
		if len(name) > 40 {
			name = name[:37] + "..."
		}
		fmt.Printf("%-8s %-40s %-12s %-5d\n", a.ID, name, a.Scores.AICapability, a.Scores.AGIWave)
	}
}

func getDomainFromID(id string) int {
	parts := strings.Split(id, ".")
	if len(parts) > 0 {
		n, _ := strconv.Atoi(parts[0])
		return n
	}
	return 0
}

func cmdActivity(id string) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var activity *Activity
	for _, a := range activities {
		if a.ID == id {
			activity = &a
			break
		}
	}
	if activity == nil {
		fmt.Fprintf(os.Stderr, "Activity %s not found\n", id)
		os.Exit(1)
	}

	fmt.Printf("Activity %s: %s\n", activity.ID, activity.Name)
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Description:  %s\n", activity.Description)
	fmt.Printf("Category:     %s\n", activity.CategoryID)
	if activity.DomainID > 0 {
		fmt.Printf("Domain:       %d\n", activity.DomainID)
	}
	fmt.Println()
	fmt.Println("Scores:")
	fmt.Printf("  Abstraction:       %d\n", activity.Scores.Abstraction)
	fmt.Printf("  Error Tolerance:   %d\n", activity.Scores.ErrorTolerance)
	fmt.Printf("  Feedback Speed:    %d\n", activity.Scores.FeedbackSpeed)
	fmt.Printf("  Social Complexity: %d\n", activity.Scores.SocialComplexity)
	fmt.Printf("  Purpose:           %d (%s)\n", activity.Scores.Purpose, getPurposeName(activity.Scores.Purpose))
	fmt.Printf("  AI Capability:     %s\n", activity.Scores.AICapability)
	fmt.Printf("  Bottleneck:        %s\n", activity.Scores.Bottleneck)
	fmt.Printf("  AGI Wave:          %d\n", activity.Scores.AGIWave)
	fmt.Println()
	if len(activity.ExampleTasks) > 0 {
		fmt.Println("Example Tasks:")
		for _, t := range activity.ExampleTasks {
			fmt.Printf("  - %s\n", t)
		}
	}
}

// getPurposeName returns the name for a purpose level
func getPurposeName(level int) string {
	names := map[int]string{
		1: "Productive",
		2: "Operational",
		3: "Developmental",
		4: "Relational",
		5: "Restorative",
	}
	if name, ok := names[level]; ok {
		return name
	}
	return "Unknown"
}

// cmdTable shows all activities with index values, grouped by domain
func cmdTable() {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tax, err := loadTaxonomy()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Group activities by domain
	byDomain := make(map[int][]Activity)
	for _, a := range activities {
		domainID := getDomainFromID(a.ID)
		byDomain[domainID] = append(byDomain[domainID], a)
	}

	// Print by domain
	for domainID := 1; domainID <= 10; domainID++ {
		domainActivities, ok := byDomain[domainID]
		if !ok || len(domainActivities) == 0 {
			continue
		}

		// Find domain name
		domainName := fmt.Sprintf("Domain %d", domainID)
		for _, d := range tax.Domains {
			if d.ID == domainID {
				domainName = d.Name
				break
			}
		}

		fmt.Printf("\n%s %d: %s %s\n", "═══", domainID, domainName, strings.Repeat("═", 60-len(domainName)))
		fmt.Printf("%-8s %-36s %5s %5s %5s\n", "ID", "Name", "Abstr", "Error", "Purp")
		fmt.Println(strings.Repeat("─", 70))

		for _, a := range domainActivities {
			name := a.Name
			if len(name) > 36 {
				name = name[:33] + "..."
			}
			fmt.Printf("%-8s %-36s %5d %5d %5d\n",
				a.ID, name,
				a.Scores.Abstraction,
				a.Scores.ErrorTolerance,
				a.Scores.Purpose)
		}
	}
	fmt.Println()
}

func cmdWave(wave int) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	timelines := map[int]string{
		1: "2024-2026",
		2: "2026-2028",
		3: "2028-2032",
		4: "2032+",
	}

	fmt.Printf("AGI Wave %d Activities (%s)\n", wave, timelines[wave])
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-12s %-10s\n", "ID", "Name", "Capability", "Bottleneck")
	fmt.Println(strings.Repeat("-", 80))

	count := 0
	for _, a := range activities {
		if a.Scores.AGIWave == wave {
			name := a.Name
			if len(name) > 40 {
				name = name[:37] + "..."
			}
			fmt.Printf("%-8s %-40s %-12s %-10s\n", a.ID, name, a.Scores.AICapability, a.Scores.Bottleneck)
			count++
		}
	}
	fmt.Printf("\nTotal: %d activities\n", count)
}

func cmdCapability(status string) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Activities with AI Capability: %s\n", status)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-10s %-5s\n", "ID", "Name", "Bottleneck", "Wave")
	fmt.Println(strings.Repeat("-", 80))

	count := 0
	for _, a := range activities {
		if a.Scores.AICapability == status {
			name := a.Name
			if len(name) > 40 {
				name = name[:37] + "..."
			}
			fmt.Printf("%-8s %-40s %-10s %-5d\n", a.ID, name, a.Scores.Bottleneck, a.Scores.AGIWave)
			count++
		}
	}
	fmt.Printf("\nTotal: %d activities\n", count)
}

func cmdBottleneck(bottleneck string) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Activities with Bottleneck: %s\n", bottleneck)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-12s %-5s\n", "ID", "Name", "Capability", "Wave")
	fmt.Println(strings.Repeat("-", 80))

	count := 0
	for _, a := range activities {
		if a.Scores.Bottleneck == bottleneck {
			name := a.Name
			if len(name) > 40 {
				name = name[:37] + "..."
			}
			fmt.Printf("%-8s %-40s %-12s %-5d\n", a.ID, name, a.Scores.AICapability, a.Scores.AGIWave)
			count++
		}
	}
	fmt.Printf("\nTotal: %d activities\n", count)
}

func cmdSearch(term string) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	term = strings.ToLower(term)
	fmt.Printf("Search results for: %s\n", term)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-12s %-5s\n", "ID", "Name", "Capability", "Wave")
	fmt.Println(strings.Repeat("-", 80))

	count := 0
	for _, a := range activities {
		if strings.Contains(strings.ToLower(a.Name), term) ||
			strings.Contains(strings.ToLower(a.Description), term) {
			name := a.Name
			if len(name) > 40 {
				name = name[:37] + "..."
			}
			fmt.Printf("%-8s %-40s %-12s %-5d\n", a.ID, name, a.Scores.AICapability, a.Scores.AGIWave)
			count++
		}
	}
	fmt.Printf("\nTotal: %d activities\n", count)
}

func cmdTime() {
	mappings, err := loadMappings()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	atus := mappings.ATUSMapping
	fmt.Println("ATUS Time-Spent Data (Average Minutes Per Day)")
	fmt.Printf("Source: %s\n", atus.DataSource)
	fmt.Println(strings.Repeat("-", 75))
	fmt.Printf("%-4s %-40s %-8s %-8s\n", "Code", "Category", "Min/Day", "Partic%")
	fmt.Println(strings.Repeat("-", 75))

	for _, e := range atus.Mappings {
		cat := e.ATUSCategory
		if len(cat) > 40 {
			cat = cat[:37] + "..."
		}
		fmt.Printf("%-4s %-40s %-8d %-8.0f%%\n", e.ATUSCode, cat, e.AvgMinutesPerDay, e.ParticipationRate*100)
	}

	fmt.Println(strings.Repeat("-", 75))
	fmt.Println("\nDaily Time Summary:")
	fmt.Printf("  Sleep & Personal Care: %d min (%.1f hrs)\n", atus.Summary.SleepAndPersonalCare, float64(atus.Summary.SleepAndPersonalCare)/60)
	fmt.Printf("  Work:                  %d min (%.1f hrs)\n", atus.Summary.Work, float64(atus.Summary.Work)/60)
	fmt.Printf("  Leisure:               %d min (%.1f hrs)\n", atus.Summary.Leisure, float64(atus.Summary.Leisure)/60)
	fmt.Printf("  Household & Care:      %d min (%.1f hrs)\n", atus.Summary.HouseholdAndCare, float64(atus.Summary.HouseholdAndCare)/60)
	fmt.Printf("  Travel:                %d min (%.1f hrs)\n", atus.Summary.Travel, float64(atus.Summary.Travel)/60)
	fmt.Printf("  Other:                 %d min (%.1f hrs)\n", atus.Summary.Other, float64(atus.Summary.Other)/60)
	fmt.Printf("  Total:                 %d min (24 hrs)\n", atus.Summary.TotalMinutesPerDay)
}

func cmdEcon() {
	mappings, err := loadMappings()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	econ := mappings.EconomicImpact
	fmt.Printf("Economic Impact by HAAI Domain (%s %d)\n", econ.Currency, econ.Year)
	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf("%-4s %-28s %12s %8s %12s %-12s\n", "ID", "Domain", "Workers", "% Work", "Value ($B)", "Automation")
	fmt.Println(strings.Repeat("-", 90))

	for _, d := range econ.DomainEcon {
		name := d.DomainName
		if len(name) > 28 {
			name = name[:25] + "..."
		}
		fmt.Printf("%-4d %-28s %12s %7.1f%% %12d %-12s\n",
			d.DomainID, name,
			formatNumber(d.EstimatedWorkers),
			d.PercentOfWorkforce,
			d.AnnualValueBillions,
			d.AutomationExposure)
	}

	fmt.Println(strings.Repeat("-", 90))
	fmt.Println("\nUS Labor Market Summary:")
	fmt.Printf("  Total Employment:     %s workers\n", formatNumber(econ.USLaborMarket.TotalEmployment))
	fmt.Printf("  Total Wages:          $%.1f trillion\n", float64(econ.USLaborMarket.TotalWages)/1e12)
	fmt.Printf("  Average Hourly Wage:  $%.2f\n", econ.USLaborMarket.AverageHourlyWage)
}

func formatNumber(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1e6)
	}
	if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1e3)
	}
	return fmt.Sprintf("%d", n)
}

func cmdStats() {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Count by capability
	capCounts := make(map[string]int)
	waveCounts := make(map[int]int)
	bottleneckCounts := make(map[string]int)
	domainCounts := make(map[int]int)
	purposeCounts := make(map[int]int)

	for _, a := range activities {
		capCounts[a.Scores.AICapability]++
		waveCounts[a.Scores.AGIWave]++
		bottleneckCounts[a.Scores.Bottleneck]++
		domainCounts[getDomainFromID(a.ID)]++
		purposeCounts[a.Scores.Purpose]++
	}

	fmt.Println("HAAI Taxonomy Statistics")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total Activities: %d\n\n", len(activities))

	fmt.Println("By AI Capability:")
	for _, cap := range []string{"solved", "near_solved", "partial", "early", "not_attempted"} {
		pct := float64(capCounts[cap]) / float64(len(activities)) * 100
		fmt.Printf("  %-15s %4d (%5.1f%%)\n", cap, capCounts[cap], pct)
	}

	fmt.Println("\nBy AGI Wave:")
	for wave := 1; wave <= 4; wave++ {
		pct := float64(waveCounts[wave]) / float64(len(activities)) * 100
		fmt.Printf("  Wave %d:         %4d (%5.1f%%)\n", wave, waveCounts[wave], pct)
	}

	fmt.Println("\nBy Purpose:")
	for purpose := 1; purpose <= 5; purpose++ {
		pct := float64(purposeCounts[purpose]) / float64(len(activities)) * 100
		fmt.Printf("  %d %-13s %4d (%5.1f%%)\n", purpose, getPurposeName(purpose), purposeCounts[purpose], pct)
	}

	fmt.Println("\nBy Bottleneck:")
	var bottlenecks []string
	for b := range bottleneckCounts {
		bottlenecks = append(bottlenecks, b)
	}
	sort.Slice(bottlenecks, func(i, j int) bool {
		return bottleneckCounts[bottlenecks[i]] > bottleneckCounts[bottlenecks[j]]
	})
	for _, b := range bottlenecks {
		if b == "" {
			continue
		}
		pct := float64(bottleneckCounts[b]) / float64(len(activities)) * 100
		fmt.Printf("  %-12s %4d (%5.1f%%)\n", b, bottleneckCounts[b], pct)
	}

	fmt.Println("\nBy Domain:")
	for d := 1; d <= 10; d++ {
		fmt.Printf("  Domain %2d:     %4d activities\n", d, domainCounts[d])
	}
}

// cmdIndex shows details about an index (abstraction, error-tolerance, or purpose)
func cmdIndex(name string) {
	indices, err := loadIndices()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading indices: %v\n", err)
		os.Exit(1)
	}

	idx, ok := indices[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Index '%s' not found. Available: abstraction, error-tolerance, purpose\n", name)
		os.Exit(1)
	}

	fmt.Printf("Index: %s\n", idx.IndexName)
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("ID:          %s\n", idx.IndexID)
	fmt.Printf("Version:     %s\n", idx.Version)
	fmt.Printf("Description: %s\n", idx.Description)
	fmt.Println()
	fmt.Println("Scale:")
	fmt.Printf("  Range: %d - %d\n", idx.Scale.Min, idx.Scale.Max)
	fmt.Println()
	fmt.Println("Levels:")
	for _, level := range idx.Scale.Levels {
		fmt.Printf("  %d. %-18s %s\n", level.Level, level.Name, level.Definition)
	}

	// Count activities per level
	fmt.Println()
	fmt.Println("Distribution:")
	levelCounts := make(map[int]int)
	for _, val := range idx.Values {
		levelCounts[val]++
	}
	total := len(idx.Values)
	for level := idx.Scale.Min; level <= idx.Scale.Max; level++ {
		count := levelCounts[level]
		pct := float64(count) / float64(total) * 100
		levelName := ""
		for _, l := range idx.Scale.Levels {
			if l.Level == level {
				levelName = l.Name
				break
			}
		}
		fmt.Printf("  Level %d %-16s %4d activities (%5.1f%%)\n", level, levelName, count, pct)
	}
}

// cmdPurpose lists activities by purpose level (1-5)
func cmdPurpose(level int) {
	activities, err := loadActivities()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	purposeName := getPurposeName(level)
	fmt.Printf("Activities with Purpose Level %d (%s)\n", level, purposeName)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-8s %-40s %-12s %-5s\n", "ID", "Name", "Capability", "Wave")
	fmt.Println(strings.Repeat("-", 80))

	count := 0
	for _, a := range activities {
		if a.Scores.Purpose == level {
			name := a.Name
			if len(name) > 40 {
				name = name[:37] + "..."
			}
			fmt.Printf("%-8s %-40s %-12s %-5d\n", a.ID, name, a.Scores.AICapability, a.Scores.AGIWave)
			count++
		}
	}
	fmt.Printf("\nTotal: %d activities\n", count)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "domains":
		cmdDomains()
	case "domain":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai domain <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid domain ID: %s\n", args[0])
			os.Exit(1)
		}
		cmdDomain(id)
	case "activities":
		domainFilter := 0
		if len(args) > 0 {
			domainFilter, _ = strconv.Atoi(args[0])
		}
		cmdActivities(domainFilter)
	case "activity":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai activity <id>")
			os.Exit(1)
		}
		cmdActivity(args[0])
	case "wave":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai wave <1-4>")
			os.Exit(1)
		}
		wave, err := strconv.Atoi(args[0])
		if err != nil || wave < 1 || wave > 4 {
			fmt.Fprintf(os.Stderr, "Invalid wave: %s (must be 1-4)\n", args[0])
			os.Exit(1)
		}
		cmdWave(wave)
	case "capability":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai capability <solved|near_solved|partial|early|not_attempted>")
			os.Exit(1)
		}
		cmdCapability(args[0])
	case "bottleneck":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai bottleneck <type>")
			os.Exit(1)
		}
		cmdBottleneck(args[0])
	case "index":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai index <name>")
			fmt.Fprintln(os.Stderr, "Available indices: abstraction, error-tolerance, purpose")
			os.Exit(1)
		}
		cmdIndex(args[0])
	case "purpose":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai purpose <1-5>")
			os.Exit(1)
		}
		level, err := strconv.Atoi(args[0])
		if err != nil || level < 1 || level > 5 {
			fmt.Fprintf(os.Stderr, "Invalid purpose level: %s (must be 1-5)\n", args[0])
			os.Exit(1)
		}
		cmdPurpose(level)
	case "search":
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Usage: haai search <term>")
			os.Exit(1)
		}
		cmdSearch(strings.Join(args, " "))
	case "time":
		cmdTime()
	case "econ":
		cmdEcon()
	case "stats":
		cmdStats()
	case "table":
		cmdTable()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}
