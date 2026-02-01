# Comprehensive Human Activity Taxonomy for AGI Progress Mapping

**Version:** 1.0 (Draft)
**Date:** 2026-02-01
**Purpose:** Predict AI automation timelines across all human activities

---

## Executive Summary

This taxonomy classifies all human activities—paid work and unpaid life—into a structure optimized for predicting when artificial intelligence will achieve human-level capability at each activity.

**Core Thesis:**
- Software AGI operates in abstract/symbolic space → solves abstract tasks first
- Embodied AGI requires physical interaction → solves physical tasks later
- Primary predictors: (1) **Abstraction level**, (2) **Error tolerance**

**Structure:**
- 10 Domains (Level 1) organized by abstraction-embodiment spectrum
- 47 Categories (Level 2) with 3-6 per domain
- ~200-400 Activities (Level 3) to be populated with scores

---

## Part A: Methodology

### 1. Design Principles

1. **Predict, don't describe:** Optimize for attributes that predict AI capability timing, not occupational classification.

2. **Task, not job:** A nurse performs dozens of tasks; each belongs in the taxonomy separately.

3. **Observable, not inferred:** Activities are objectively identifiable, not based on mental states.

4. **Technology-agnostic:** Define by human capability required, not current technology.

5. **Globally applicable:** Avoid Western/developed-world bias.

6. **MECE compliant:** Mutually exclusive, collectively exhaustive—every activity fits exactly one category.

### 2. Organizing Principle: Abstraction-Embodiment Spectrum

Domains are ordered from **highest abstraction** (earliest AI automation) to **lowest abstraction** (latest AI automation):

| Domain | Abstraction Score | Primary AI System Type |
|--------|-------------------|------------------------|
| 1. Symbolic Computation | 10/10 | Rule engines, calculators |
| 2. Information Synthesis | 9/10 | Language models, RAG |
| 3. Creative Generation | 8/10 | Generative AI |
| 4. Interpersonal Communication | 7/10 | Conversational AI |
| 5. Procedural Service Delivery | 6/10 | Multimodal AI + workflows |
| 6. Care & Relational Support | 5/10 | Affective computing |
| 7. Structured Physical Operations | 4/10 | Industrial robotics |
| 8. Skilled Physical Craft | 3/10 | Dexterous manipulation |
| 9. Environmental Navigation | 2/10 | Mobile robotics, AV |
| 10. Personal Embodiment | 1/10 | General-purpose home robots |

### 3. Rationale for This Structure

**Why abstraction predicts automation:**
- High-abstraction activities (symbolic, informational) require only software
- Low-abstraction activities require physical sensing, planning, and actuation
- The software → hardware progression maps directly to AI development trajectory

**Why error tolerance is secondary:**
- Within any domain, low-error-tolerance activities automate later
- Error tolerance interacts with abstraction (physical errors often irreversible)
- Captures the "AI-in-the-loop vs. human-in-the-loop" dimension

---

## Part B: Attribute Definitions

### Attribute 1: Abstraction Level (1-5 Scale)

| Level | Name | Definition | Signal |
|-------|------|------------|--------|
| 1 | Pure Abstract | Operates entirely in symbol space. No physical grounding required. | Input/output are text, numbers, code |
| 2 | Grounded Abstract | Reasons about physical world but acts through digital intermediaries. | Interprets sensor data, outputs decisions |
| 3 | Supervised Physical | Physical action in structured, predictable environments. | Repeatable motions, known objects |
| 4 | Adaptive Physical | Physical action requiring real-time adaptation to variability. | Novel objects, changing conditions |
| 5 | Deeply Embodied | Full sensorimotor integration, social presence, human-like dexterity. | Touch sensitivity, emotional attunement |

### Attribute 2: Error Tolerance (1-4 Scale)

| Level | Name | Definition | Examples |
|-------|------|------------|----------|
| 1 | High | Errors easily corrected, low stakes, iteration expected. | Drafting, brainstorming, prototyping |
| 2 | Medium | Errors costly but recoverable, moderate stakes. | Cooking, scheduling, routine maintenance |
| 3 | Low | Errors expensive or harmful, limited recovery. | Surgery, vehicle repair, legal filing |
| 4 | Very Low | Errors potentially catastrophic or irreversible. | Emergency response, childcare safety |

### Attribute 3: Feedback Loop Speed (1-4 Scale)

| Level | Timeframe | Examples |
|-------|-----------|----------|
| 1 | Seconds to minutes | Conversation, object manipulation |
| 2 | Minutes to hours | Task completion, cooking |
| 3 | Hours to days | Project work, analysis |
| 4 | Days to months | Strategy, research |

### Attribute 4: Social Complexity (0-4 Scale)

| Level | Name | Definition |
|-------|------|------------|
| 0 | None | Solo task, no human interaction |
| 1 | Informational | Exchange information with humans |
| 2 | Coordinative | Synchronize actions with humans |
| 3 | Persuasive | Influence human beliefs or behavior |
| 4 | Relational | Build trust, provide emotional support |

### Attribute 5: Current AI Capability (Categorical)

| Status | Definition |
|--------|------------|
| Solved | AI matches or exceeds median human performance |
| Near-solved | AI approaches human performance, deployment underway |
| Partial | AI assists but cannot fully replace humans |
| Early | Research prototypes only |
| Not attempted | No serious AI efforts yet |

### Attribute 6: Primary Bottleneck (Categorical)

| Bottleneck | Description |
|------------|-------------|
| None | Solved |
| Sensing | Perception limitations |
| Reasoning | Cognitive complexity |
| Dexterity | Fine motor control |
| Mobility | Locomotion, navigation |
| Adaptation | Handling novel situations |
| Social | Human interaction requirements |
| Safety | Risk tolerance requirements |
| Regulation | Legal/policy barriers |
| Data | Insufficient training data |

### Attribute 7: Estimated AGI Wave (1-4)

| Wave | Timeline | Characteristics |
|------|----------|-----------------|
| 1 | 2024-2026 | Current LLM capabilities + basic robotics |
| 2 | 2026-2028 | Multimodal AI + structured environments |
| 3 | 2028-2032 | Adaptive embodied AI + social modeling |
| 4 | 2032+ | General-purpose embodied AI |

---

## Part C: Domain Definitions (Level 1)

### Domain 1: SYMBOLIC COMPUTATION

**Definition:** Pure transformation of formal symbols, numbers, or structured data without reference to real-world meaning.

**Abstraction Score:** 10/10
**Typical Error Tolerance:** Variable (depends on stakes)
**Estimated AGI Wave:** 1 (largely automated)

**Inclusion Criteria:**
- Input/output are formal symbols (numbers, code, logical expressions)
- Transformation rules are explicit and deterministic
- No interpretation of real-world meaning required
- Correctness is objectively verifiable

**Exclusion Criteria:**
- Requires understanding natural language semantics
- Requires world knowledge for interpretation
- Involves judgment about edge cases

---

### Domain 2: INFORMATION SYNTHESIS

**Definition:** Combining, analyzing, and transforming information from multiple sources to produce new informational outputs.

**Abstraction Score:** 9/10
**Typical Error Tolerance:** Medium-High
**Estimated AGI Wave:** 1-2

**Inclusion Criteria:**
- Primary input/output are information (text, data, documents)
- Requires synthesis across multiple sources
- Output evaluated on informational quality
- Can be performed entirely through digital interfaces

**Exclusion Criteria:**
- Requires physical presence or manipulation
- Primary output is a physical artifact
- Requires real-time social interaction

---

### Domain 3: CREATIVE GENERATION

**Definition:** Producing novel artifacts (text, images, code, music, designs) that did not exist before.

**Abstraction Score:** 8/10
**Typical Error Tolerance:** High (iteration expected)
**Estimated AGI Wave:** 1-2

**Inclusion Criteria:**
- Output is a novel artifact
- Success measured by creativity, aesthetics, or functional innovation
- Requires generation rather than transformation
- Output exists in digital or symbolic form

**Exclusion Criteria:**
- Pure reproduction or copying
- Output is physical (see Domain 8)
- No creative latitude allowed

---

### Domain 4: INTERPERSONAL COMMUNICATION

**Definition:** Exchanging information, ideas, and intentions between humans through primarily digital or verbal channels.

**Abstraction Score:** 7/10
**Typical Error Tolerance:** Medium
**Estimated AGI Wave:** 2

**Inclusion Criteria:**
- Primary purpose is information exchange between people
- Success depends on mutual understanding
- Can be conducted via digital channels
- Involves interpretation of intent and context

**Exclusion Criteria:**
- Physical presence essential (see Domain 6)
- Primarily about emotional support (see Domain 6)
- Purely transactional (see Domain 5)

---

### Domain 5: PROCEDURAL SERVICE DELIVERY

**Definition:** Providing services to others by following established procedures and protocols.

**Abstraction Score:** 6/10
**Typical Error Tolerance:** Medium
**Estimated AGI Wave:** 2-3

**Inclusion Criteria:**
- Follows defined procedures or protocols
- Service delivered to customers/clients/citizens
- Success measured by procedure completion
- Interaction is transactional rather than relational

**Exclusion Criteria:**
- Requires deep personal relationship (see Domain 6)
- Primarily physical manipulation (see Domain 7-8)
- High creative latitude (see Domain 3)

---

### Domain 6: CARE & RELATIONAL SUPPORT

**Definition:** Activities primarily focused on human wellbeing, emotional support, and relationship maintenance.

**Abstraction Score:** 5/10
**Typical Error Tolerance:** Low (trust is fragile)
**Estimated AGI Wave:** 3-4

**Inclusion Criteria:**
- Primary purpose is human wellbeing or relationship
- Success depends on trust, empathy, or emotional attunement
- Human presence has intrinsic value
- Involves vulnerability of care recipient

**Exclusion Criteria:**
- Purely informational exchange (see Domain 4)
- Transactional service (see Domain 5)
- Can be fully automated without loss of value

---

### Domain 7: STRUCTURED PHYSICAL OPERATIONS

**Definition:** Physical manipulation and movement within controlled, predictable environments.

**Abstraction Score:** 4/10
**Typical Error Tolerance:** Variable (industry-dependent)
**Estimated AGI Wave:** 2-3

**Inclusion Criteria:**
- Primary activity involves physical manipulation or operation
- Environment is controlled and predictable
- Procedures are standardized
- Physical objects are uniform or classified

**Exclusion Criteria:**
- Environment is unstructured (see Domain 9)
- Requires high dexterity/craft skill (see Domain 8)
- Primarily about human interaction (see Domain 4-6)

---

### Domain 8: SKILLED PHYSICAL CRAFT

**Definition:** Physical manipulation requiring expertise, dexterity, and judgment.

**Abstraction Score:** 3/10
**Typical Error Tolerance:** Low-Medium
**Estimated AGI Wave:** 3-4

**Inclusion Criteria:**
- Requires acquired physical skill and expertise
- Conditions vary requiring adaptation
- Quality depends on craftsperson judgment
- Cannot be fully procedualized

**Exclusion Criteria:**
- Fully standardized operations (see Domain 7)
- Primarily about locomotion (see Domain 9)
- Purely symbolic output (see Domain 1-3)

---

### Domain 9: ENVIRONMENTAL NAVIGATION

**Definition:** Moving through and interacting with unstructured, dynamic environments.

**Abstraction Score:** 2/10
**Typical Error Tolerance:** Low (physical safety)
**Estimated AGI Wave:** 2-4

**Inclusion Criteria:**
- Primary activity is locomotion or navigation
- Environment is unstructured or dynamic
- Requires real-time perception and decision-making
- Physical safety is a primary constraint

**Exclusion Criteria:**
- Controlled environment (see Domain 7)
- Stationary manipulation (see Domain 8)
- Social navigation (see Domain 4-6)

---

### Domain 10: PERSONAL EMBODIMENT

**Definition:** Activities for maintaining one's own body, meeting basic human needs, and personal self-expression through physical activity.

**Abstraction Score:** 1/10
**Typical Error Tolerance:** Variable
**Estimated AGI Wave:** 3-4 (partial assistance)

**Inclusion Criteria:**
- Activity is self-directed (for oneself)
- Involves one's own body
- Fundamental human need or expression
- Cannot be delegated without losing essential meaning

**Exclusion Criteria:**
- Service provided to others (see Domain 5-6)
- Work performed for compensation (see other domains)
- Can be meaningfully automated for the person

---

## Part D: Category Definitions (Level 2)

### Domain 1: SYMBOLIC COMPUTATION (4 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 1.1 | Numerical Computation | Arithmetic, statistical calculations, spreadsheet operations |
| 1.2 | Data Transformation | ETL, format conversion, database operations |
| 1.3 | Formal Verification | Code compilation, syntax checking, validation |
| 1.4 | Algorithmic Execution | Running defined procedures, sorting, searching |

---

### Domain 2: INFORMATION SYNTHESIS (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 2.1 | Research & Analysis | Literature review, market research, data analysis |
| 2.2 | Report Generation | Writing reports, summaries, documentation |
| 2.3 | Information Curation | Organizing, tagging, categorizing information |
| 2.4 | Pattern Recognition | Identifying trends, anomalies, correlations |
| 2.5 | Recommendation Systems | Evaluating options, producing ranked suggestions |

---

### Domain 3: CREATIVE GENERATION (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 3.1 | Textual Creation | Writing fiction, copy, scripts, articles |
| 3.2 | Visual Design | Graphic design, UI/UX, illustration |
| 3.3 | Software Development | Writing code, system architecture, algorithms |
| 3.4 | Musical Composition | Creating music, sound design, audio production |
| 3.5 | Strategic Innovation | Business strategy, product design, invention |

---

### Domain 4: INTERPERSONAL COMMUNICATION (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 4.1 | Instructional Communication | Teaching, training, explaining |
| 4.2 | Persuasive Communication | Sales, negotiation, advocacy |
| 4.3 | Coordinative Communication | Meetings, project management, scheduling |
| 4.4 | Informational Exchange | Customer service, technical support, Q&A |
| 4.5 | Social Communication | Casual conversation, networking, relationship maintenance |

---

### Domain 5: PROCEDURAL SERVICE DELIVERY (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 5.1 | Administrative Processing | Forms, applications, records management |
| 5.2 | Retail & Hospitality | Sales transactions, food service, accommodation |
| 5.3 | Financial Services | Banking, insurance, accounting services |
| 5.4 | Healthcare Administration | Scheduling, billing, records, triage |
| 5.5 | Legal & Compliance Services | Document preparation, compliance checking |

---

### Domain 6: CARE & RELATIONAL SUPPORT (6 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 6.1 | Childcare & Development | Parenting, early childhood education, supervision |
| 6.2 | Eldercare & Disability Support | Assistance with daily living, companionship |
| 6.3 | Healthcare Delivery | Nursing, therapy, medical care (direct patient contact) |
| 6.4 | Psychological Support | Counseling, coaching, emotional support |
| 6.5 | Community & Civic Care | Volunteering, community organizing, religious service |
| 6.6 | Social Relationship Maintenance | Family time, friendship, intimacy |

---

### Domain 7: STRUCTURED PHYSICAL OPERATIONS (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 7.1 | Manufacturing & Assembly | Production lines, assembly, packaging |
| 7.2 | Warehouse & Logistics | Picking, packing, inventory management |
| 7.3 | Machine Operation | Operating industrial equipment, CNC, printing |
| 7.4 | Quality Control & Inspection | Testing, measuring, inspecting products |
| 7.5 | Facility Maintenance | Cleaning, maintenance of commercial spaces |

---

### Domain 8: SKILLED PHYSICAL CRAFT (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 8.1 | Construction & Building | Carpentry, plumbing, electrical, masonry |
| 8.2 | Equipment Repair & Maintenance | Automotive, HVAC, appliance repair |
| 8.3 | Surgical & Medical Procedures | Surgery, dental procedures, medical interventions |
| 8.4 | Culinary Arts | Professional cooking, food preparation |
| 8.5 | Artisan Crafts | Woodworking, metalwork, textiles, ceramics |

---

### Domain 9: ENVIRONMENTAL NAVIGATION (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 9.1 | Vehicle Operation | Driving cars, trucks, forklifts |
| 9.2 | Last-Mile Delivery | Package delivery, food delivery |
| 9.3 | Field Work | Agriculture, surveying, outdoor maintenance |
| 9.4 | Exploration & Search | Search and rescue, inspection of novel environments |
| 9.5 | Recreational Navigation | Hiking, sports, outdoor recreation |

---

### Domain 10: PERSONAL EMBODIMENT (5 Categories)

| ID | Category | Description |
|----|----------|-------------|
| 10.1 | Basic Self-Care | Sleeping, bathing, grooming, dressing |
| 10.2 | Nourishment | Eating, meal preparation for self |
| 10.3 | Health Maintenance | Exercise, medical self-care, medication |
| 10.4 | Domestic Maintenance | Cleaning, laundry, home organization |
| 10.5 | Personal Expression | Hobbies, leisure activities, rest |

---

## Part E: Mapping Appendix

### E.1: O*NET Work Activities Mapping

| O*NET Category | O*NET Activities | Taxonomy Domain(s) |
|----------------|------------------|-------------------|
| **Information Input** | Getting Information, Monitoring, Identifying Objects, Inspecting, Estimating | 1, 2, 7.4 |
| **Mental Processes** | Processing Info, Evaluating Compliance, Analyzing Data, Judging Quality, Decision Making, Creative Thinking, Updating Knowledge, Strategy Development, Scheduling, Organizing | 1, 2, 3, 4.3 |
| **Work Output (Digital)** | Working with Computers, Documenting Information | 1, 2, 3 |
| **Work Output (Technical)** | Drafting/Specifying, Repairing Mechanical/Electronic | 3.3, 8.1, 8.2 |
| **Work Output (Physical)** | Physical Activities, Handling Objects, Controlling Machines, Operating Vehicles | 7, 8, 9 |
| **Interacting (Communication)** | Interpreting, Communicating (internal/external) | 4 |
| **Interacting (Relationships)** | Establishing Relationships, Caring, Selling, Negotiating | 4, 5, 6 |
| **Interacting (Leadership)** | Coordinating, Team Building, Training, Coaching, Consulting | 4.1, 4.3, 6.4 |
| **Interacting (Admin)** | Performing Admin, Staffing, Controlling Resources | 5.1, 5.3 |

### E.2: ATUS Categories Mapping

| ATUS Code | ATUS Category | Taxonomy Domain(s) |
|-----------|---------------|-------------------|
| 01 | Personal Care | 10.1, 10.3 |
| 02 | Household Activities | 10.2, 10.4 |
| 03 | Caring for Household Members | 6.1, 6.2 |
| 04 | Caring for Non-Household Members | 6.1, 6.2, 6.5 |
| 05 | Work and Work-Related | 1-8 (by occupation) |
| 06 | Education | 2.1, 4.1 |
| 07 | Consumer Purchases | 5.2, 9.2 |
| 08 | Professional and Personal Care Services | 5, 6.3, 6.4 |
| 09 | Household Services | 5.1 |
| 10 | Government Services and Civic | 5.5, 6.5 |
| 11 | Eating and Drinking | 10.2 |
| 12 | Socializing, Relaxing, and Leisure | 4.5, 6.6, 10.5 |
| 13 | Sports, Exercise, and Recreation | 9.5, 10.3 |
| 14 | Religious and Spiritual | 6.5 |
| 15 | Volunteer Activities | 6.5 |
| 16 | Telephone Calls | 4 |
| 18 | Traveling | 9.1, 9.2 |

### E.3: BEHAVIOR-1K Task Categories Mapping

| BEHAVIOR Category | Taxonomy Domain(s) |
|-------------------|-------------------|
| Cleaning/Wiping | 7.5, 10.4 |
| Cooking/Freezing | 8.4, 10.2 |
| Rearrangement | 7.2, 10.4 |
| Slicing/Dicing | 8.4 |
| Baking | 8.4 |
| Laundry | 7.5, 10.4 |
| Painting/Spraying | 8.1, 8.5 |
| Hanging/Installing | 8.1 |

### E.4: Kinetics/ActivityNet Action Classes Mapping

| Action Type | Taxonomy Domain(s) |
|-------------|-------------------|
| Human-Object Interaction | 7, 8, 10 |
| Body-Motion Only | 9.5, 10.3 |
| Human-Human Interaction | 4, 6 |
| Playing Instruments | 3.4, 10.5 |
| Sports | 9.5, 10.3 |

---

## Part F: Validation Checklist

### F.1: Day-in-the-Life Coverage Test

| Time | Activity | Taxonomy Classification |
|------|----------|------------------------|
| 6:00 | Wake up, use bathroom | 10.1 Basic Self-Care |
| 6:15 | Shower, dress | 10.1 Basic Self-Care |
| 6:30 | Prepare breakfast | 10.2 Nourishment |
| 7:00 | Eat breakfast | 10.2 Nourishment |
| 7:15 | Check email | 2.3 Information Curation |
| 7:30 | Commute to work | 9.1 Vehicle Operation |
| 8:00 | Team standup meeting | 4.3 Coordinative Communication |
| 8:30 | Code review | 3.3 Software Development |
| 10:00 | Debug software issue | 3.3 Software Development |
| 12:00 | Lunch with colleague | 4.5 Social Communication |
| 13:00 | Write technical documentation | 2.2 Report Generation |
| 15:00 | Client presentation | 4.2 Persuasive Communication |
| 17:00 | Commute home | 9.1 Vehicle Operation |
| 17:30 | Exercise/run | 10.3 Health Maintenance |
| 18:30 | Prepare dinner | 10.2 Nourishment |
| 19:00 | Eat dinner with family | 6.6 Social Relationship Maintenance |
| 20:00 | Help child with homework | 6.1 Childcare & Development |
| 21:00 | Watch TV/relax | 10.5 Personal Expression |
| 22:00 | Personal hygiene, sleep | 10.1 Basic Self-Care |

**Result:** All activities classified. No gaps identified.

### F.2: Occupation Coverage Test

| ISCO Major Group | Sample Tasks | Taxonomy Coverage |
|------------------|--------------|-------------------|
| Managers | Strategy, coordination, decisions | 3.5, 4.3, 5.1 |
| Professionals | Analysis, design, consultation | 2, 3, 4.1, 6.3-6.4 |
| Technicians | Technical support, operation | 3.3, 7.3, 8.2 |
| Clerical Support | Admin, data entry, filing | 1.2, 2.3, 5.1 |
| Service and Sales | Customer service, retail | 4.4, 5.2 |
| Skilled Agricultural | Farming, animal husbandry | 9.3 |
| Craft and Trade | Construction, repair | 8.1, 8.2, 8.5 |
| Plant and Machine Operators | Manufacturing, driving | 7.1, 7.3, 9.1 |
| Elementary Occupations | Cleaning, loading, basic tasks | 7.2, 7.5, 9.2 |

**Result:** All ISCO major groups map to taxonomy categories.

### F.3: AI Capability Alignment Test

| AI System | Demonstrated Capability | Taxonomy Prediction |
|-----------|------------------------|---------------------|
| GPT-4/Claude | Text generation, analysis, code | 1-3: Wave 1-2 |
| GitHub Copilot | Code completion | 3.3: Wave 1-2 |
| Midjourney/DALL-E | Image generation | 3.2: Wave 1-2 |
| Tesla Autopilot | Highway driving | 9.1: Wave 2-3 |
| Waymo | Urban autonomous driving | 9.1: Wave 2-3 |
| Boston Dynamics Spot | Navigation, inspection | 9.4: Wave 2-3 |
| Amazon Kiva | Warehouse operations | 7.2: Wave 2 |
| Da Vinci Surgical | Assisted surgery | 8.3: Wave 3-4 |

**Result:** Current AI capabilities align with taxonomy's abstraction-based predictions.

---

## Part G: Edge Cases and Classification Rules

### G.1: Activities Spanning Multiple Domains

When an activity involves multiple domains, classify by **primary purpose**:

| Activity | Appears to be... | Classify as... | Rationale |
|----------|------------------|----------------|-----------|
| Video call with client | Communication + Service | 4.4 Informational Exchange | Primary purpose is communication |
| Writing code while pair programming | Creative + Social | 3.3 Software Development | Primary output is code |
| Teaching surgery | Education + Craft | 4.1 Instructional Communication | Primary purpose is teaching |
| Telehealth consultation | Healthcare + Communication | 6.3 Healthcare Delivery | Primary purpose is care |

### G.2: Boundary Rules

**Domain 1 vs 2:** If requires interpretation of meaning → Domain 2
**Domain 2 vs 3:** If output must be novel/creative → Domain 3
**Domain 4 vs 5:** If relationship > transaction → Domain 4; if procedure > relationship → Domain 5
**Domain 5 vs 6:** If trust/emotional attunement essential → Domain 6
**Domain 7 vs 8:** If requires acquired expertise/judgment → Domain 8
**Domain 8 vs 9:** If primary activity is locomotion → Domain 9
**Domain 9 vs 10:** If for self vs. for work → Domain 10 (self) or 9 (work)

### G.3: Context Modifiers

The same physical action may classify differently based on context:

| Physical Action | Professional Context | Personal Context |
|-----------------|---------------------|------------------|
| Cooking | 8.4 Culinary Arts | 10.2 Nourishment |
| Driving | 9.1 Vehicle Operation | 9.1 Vehicle Operation (same) |
| Cleaning | 7.5 Facility Maintenance | 10.4 Domestic Maintenance |
| Exercise | 6.3 Healthcare Delivery (PT) | 10.3 Health Maintenance |

---

## Part H: Next Steps (Level 3 Population)

### H.1: Level 3 Activity Population Approach

For each Level 2 category, enumerate 5-15 specific activities with all attributes scored. Target: 200-400 total activities.

### H.2: Scoring Methodology

Each activity will be scored by:
1. Domain experts (for current AI capability)
2. AI researchers (for bottleneck identification)
3. Time-use researchers (for validation against ATUS/HETUS)

### H.3: Validation Protocol

1. **Inter-rater reliability:** Two independent raters should achieve κ > 0.7
2. **Coverage audit:** Walk through 10 diverse occupations, ensure all tasks map
3. **Predictive validation:** Compare predictions to actual AI deployment timelines

---

## References

### Primary Sources
- O*NET Content Model: https://www.onetcenter.org/content.html
- ATUS Activity Lexicon: https://www.bls.gov/tus/lexicons.htm
- HETUS Guidelines: https://ec.europa.eu/eurostat/web/time-use-surveys
- ICATUS 2016: https://unstats.un.org/unsd/gender/timeuse/

### AI Benchmarks
- BEHAVIOR-1K: https://behavior.stanford.edu
- ActivityNet: http://activity-net.org/
- Kinetics: https://github.com/cvdfoundation/kinetics-dataset

### Cognitive Frameworks
- Bloom's Taxonomy (Cognitive Domain)
- Dave's Taxonomy (Psychomotor Domain)
- Krathwohl's Taxonomy (Affective Domain)

---

*Document Status: Level 1-2 complete. Level 3 activity enumeration pending.*
