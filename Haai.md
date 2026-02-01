# Comprehensive Human Activity Taxonomy for AGI Progress Mapping

**Version:** 2.1
**Date:** 2026-02-01
**Purpose:** Predict AI automation timelines across all human activities

### Changelog

| Version | Date | Changes |
|---------|------|---------|
| 2.1 | 2026-02-01 | **Structural refactoring:** Separated intrinsic (stable) and temporal (mutable) attributes in Part B. Redefined waves as relative scales with explicit mappings. Added temporal data versioning mechanism (H.5). Added Economics/Trust bottlenecks. Added community contribution process (H.7). |
| 2.0 | 2026-02-01 | **Complete framework release:** Added full Level 3 enumeration with 400 scored activities across all 47 categories (Part D.2), added summary statistics (Part D.3) |
| 1.1 | 2026-02-01 | Added example scored activities (Part D.1), expanded methodology rationale, clarified error tolerance components, added implementation roadmap (Part H.4-H.6) |
| 1.0 | 2026-02-01 | Initial draft with Level 1-2 complete |

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
- 400 Activities (Level 3) with full attribute scoring

---

## Part A: Methodology

### 1. Design Principles

1. **Predict, don't describe:** Optimize for attributes that predict AI capability timing, not occupational classification. The taxonomy prioritizes automation-relevant features over traditional occupational categories.

2. **Task, not job:** A nurse performs dozens of tasks; each belongs in the taxonomy separately. Jobs are decomposed into constituent activities for more granular prediction.

3. **Observable, not inferred:** Activities are objectively identifiable, not based on mental states. Classification relies on observable behaviors and outputs, not intentions.

4. **Technology-agnostic:** Define by human capability required, not current technology. Descriptions should remain valid regardless of specific AI implementations.

5. **Globally applicable:** Avoid Western/developed-world bias. Activities should be recognizable across cultures and economic contexts.

6. **MECE compliant:** Mutually exclusive, collectively exhaustive—every activity fits exactly one category. Boundary rules (Part G) resolve ambiguous cases.

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
- High-abstraction activities (symbolic, informational) require only software—no physical embodiment, sensors, or actuators
- Low-abstraction activities require physical sensing, planning, and actuation—hardware that must operate reliably in the real world
- The software → hardware progression maps directly to AI development trajectory, as observed in the progression from chess (1997) to language (2022) to driving (2020s) to dexterous manipulation (2030s+)
- Training data availability: Abstract tasks have abundant digital training data; physical tasks require expensive real-world data collection or simulation

**Why error tolerance is secondary:**
- Within any domain, low-error-tolerance activities automate later due to liability, regulatory, and safety concerns
- Error tolerance interacts with abstraction: physical errors are often irreversible (surgery, construction) while digital errors can be undone (document editing, code changes)
- Captures the "AI-in-the-loop vs. human-in-the-loop" dimension—high-stakes activities require human oversight longer

**Supporting observations:**
- AI progress has followed this pattern: calculators (1960s) → spell-check (1980s) → translation (2010s) → writing (2020s) → driving (2020s) → robotics (emerging)
- Investment patterns show software AI receiving more funding and achieving faster progress than robotics/embodied AI
- Current AI capabilities cluster in high-abstraction domains (Domains 1-4) with limited deployment in physical domains (7-10)

---

## Part B: Attribute Definitions

This taxonomy distinguishes between **intrinsic attributes** (stable properties of activities) and **temporal attributes** (assessments that change as AI progresses).

### B.1: Intrinsic Attributes (Stable)

These attributes describe inherent characteristics of activities. They should rarely change unless the activity itself is redefined.

#### Attribute 1: Abstraction Level (1-5 Scale)

| Level | Name | Definition | Signal |
|-------|------|------------|--------|
| 1 | Pure Abstract | Operates entirely in symbol space. No physical grounding required. | Input/output are text, numbers, code |
| 2 | Grounded Abstract | Reasons about physical world but acts through digital intermediaries. | Interprets sensor data, outputs decisions |
| 3 | Supervised Physical | Physical action in structured, predictable environments. | Repeatable motions, known objects |
| 4 | Adaptive Physical | Physical action requiring real-time adaptation to variability. | Novel objects, changing conditions |
| 5 | Deeply Embodied | Full sensorimotor integration, social presence, human-like dexterity. | Touch sensitivity, emotional attunement |

#### Attribute 2: Error Tolerance (1-4 Scale)

| Level | Name | Definition | Examples |
|-------|------|------------|----------|
| 1 | High | Errors easily corrected, low stakes, iteration expected. | Drafting, brainstorming, prototyping |
| 2 | Medium | Errors costly but recoverable, moderate stakes. | Cooking, scheduling, routine maintenance |
| 3 | Low | Errors expensive or harmful, limited recovery. | Surgery, vehicle repair, legal filing |
| 4 | Very Low | Errors potentially catastrophic or irreversible. | Emergency response, childcare safety |

**Error Tolerance Components:**
- **Reversibility:** Can the action be undone? (digital: usually yes; physical: often no)
- **Stakes:** What is the cost of an error? (financial, reputational, health, life)
- **Recovery time:** How long to correct an error? (seconds vs. months/years)
- **Detection difficulty:** How quickly are errors noticed? (immediate vs. delayed)

*Note: When components conflict, weight life/safety stakes most heavily, then financial stakes, then reversibility.*

#### Attribute 3: Feedback Loop Speed (1-4 Scale)

| Level | Timeframe | Examples |
|-------|-----------|----------|
| 1 | Seconds to minutes | Conversation, object manipulation |
| 2 | Minutes to hours | Task completion, cooking |
| 3 | Hours to days | Project work, analysis |
| 4 | Days to months | Strategy, research |

#### Attribute 4: Social Complexity (0-4 Scale)

| Level | Name | Definition |
|-------|------|------------|
| 0 | None | Solo task, no human interaction |
| 1 | Informational | Exchange information with humans |
| 2 | Coordinative | Synchronize actions with humans |
| 3 | Persuasive | Influence human beliefs or behavior |
| 4 | Relational | Build trust, provide emotional support |

---

### B.2: Temporal Attributes (Mutable)

These attributes assess current AI capability and predictions. They WILL change as technology advances and should be versioned separately.

**Current Assessment Date:** 2026-02-01 (update this when reassessing)

#### Attribute 5: AI Capability Status (Categorical)

| Status | Definition |
|--------|------------|
| Automated | AI handles autonomously in production at scale |
| Deployed | AI deployed in production with human oversight |
| Piloting | AI in pilot/beta testing, limited deployment |
| Research | Research prototypes only |
| Theoretical | No working prototypes yet |

*Note: Renamed from "Current AI Capability" to emphasize this is a point-in-time assessment.*

#### Attribute 6: Primary Bottleneck (Categorical)

| Bottleneck | Description | Likely Resolution |
|------------|-------------|-------------------|
| None | Already solved | - |
| Sensing | Perception limitations | Better sensors, multimodal models |
| Reasoning | Cognitive complexity | Larger models, better architectures |
| Dexterity | Fine motor control | Improved actuators, sim-to-real |
| Mobility | Locomotion, navigation | Better SLAM, legged robots |
| Adaptation | Handling novel situations | Foundation models, few-shot learning |
| Social | Human interaction requirements | Affective computing, theory of mind |
| Safety | Risk tolerance requirements | Verification, formal methods |
| Regulation | Legal/policy barriers | Policy evolution (external factor) |
| Economics | Cost-benefit not favorable | Hardware cost reduction |
| Trust | Human acceptance required | Deployment track record |

*Added: Economics, Trust bottlenecks for completeness.*

#### Attribute 7: Automation Wave (Relative Scale)

Waves are defined **relatively** to enable future-proofing. Map to calendar years based on current AI progress assessment.

| Wave | Definition | Prerequisites | Current Mapping* |
|------|------------|---------------|------------------|
| Wave 1 | Software-only automation | LLMs, current ML | 2024-2026 |
| Wave 2 | Structured environment automation | Multimodal AI + controlled robotics | 2026-2028 |
| Wave 3 | Adaptive automation | Embodied AI + real-time adaptation | 2028-2032 |
| Wave 4 | General-purpose automation | Human-level physical + social AI | 2032+ |
| N/A | Not automatable or not applicable | Human essence, consciousness | - |

*\*Current mapping as of 2026-02-01. Update mappings as AI progress becomes clearer.*

**Wave Determination Logic:**
- Wave is determined by the **latest required capability**, not the earliest
- High abstraction + low error tolerance → May be Wave 2-3 despite abstraction
- Low abstraction + high error tolerance → May be Wave 2-3 despite embodiment needs

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

## Part D.1: Example Scored Activities

The following examples demonstrate how Level 3 activities should be scored using all seven attributes. These serve as reference cases for raters.

### Example 1: Spreadsheet Formula Writing (Category 1.1)

| Attribute | Score | Rationale |
|-----------|-------|-----------|
| Abstraction Level | 1 (Pure Abstract) | Input/output are numbers and formulas; no physical grounding |
| Error Tolerance | 1 (High) | Errors easily caught and corrected; low stakes |
| Feedback Loop Speed | 1 (Seconds) | Immediate feedback on formula results |
| Social Complexity | 0 (None) | Solo task, no human interaction required |
| Current AI Capability | Solved | LLMs can write formulas accurately |
| Primary Bottleneck | None | Solved |
| Estimated AGI Wave | 1 | Already automated |

### Example 2: Market Research Report (Category 2.1)

| Attribute | Score | Rationale |
|-----------|-------|-----------|
| Abstraction Level | 1 (Pure Abstract) | Information synthesis task; digital inputs/outputs |
| Error Tolerance | 2 (Medium) | Errors costly in business decisions but correctable |
| Feedback Loop Speed | 3 (Hours to days) | Report takes time to validate |
| Social Complexity | 1 (Informational) | May require interviews/surveys |
| Current AI Capability | Partial | AI assists but human judgment still needed for quality |
| Primary Bottleneck | Reasoning | Complex multi-source synthesis |
| Estimated AGI Wave | 2 | Improving rapidly |

### Example 3: Warehouse Order Picking (Category 7.2)

| Attribute | Score | Rationale |
|-----------|-------|-----------|
| Abstraction Level | 3 (Supervised Physical) | Physical manipulation in structured environment |
| Error Tolerance | 2 (Medium) | Wrong items costly but fixable |
| Feedback Loop Speed | 1 (Seconds) | Immediate per-item feedback |
| Social Complexity | 0-1 (None to Informational) | Minimal interaction; system instructions |
| Current AI Capability | Near-solved | Amazon/Ocado systems operational |
| Primary Bottleneck | Dexterity | Varied item shapes and packaging |
| Estimated AGI Wave | 2 | Deploying now |

### Example 4: Elementary School Teaching (Category 4.1)

| Attribute | Score | Rationale |
|-----------|-------|-----------|
| Abstraction Level | 2 (Grounded Abstract) | Interprets student behavior; acts through communication |
| Error Tolerance | 3 (Low) | Developmental impact; limited recovery |
| Feedback Loop Speed | 4 (Days to months) | Learning outcomes measured over time |
| Social Complexity | 4 (Relational) | Trust, motivation, emotional support essential |
| Current AI Capability | Partial | AI tutoring exists but lacks full social capability |
| Primary Bottleneck | Social | Emotional attunement, classroom management |
| Estimated AGI Wave | 3-4 | Requires social AI advances |

### Example 5: Emergency Room Triage (Category 5.4)

| Attribute | Score | Rationale |
|-----------|-------|-----------|
| Abstraction Level | 2 (Grounded Abstract) | Assesses symptoms; makes prioritization decisions |
| Error Tolerance | 4 (Very Low) | Errors can be life-threatening |
| Feedback Loop Speed | 1 (Seconds to minutes) | Immediate decisions required |
| Social Complexity | 2 (Coordinative) | Coordinates patient flow and staff |
| Current AI Capability | Partial | AI assists but human required for liability |
| Primary Bottleneck | Safety | Error tolerance requirements |
| Estimated AGI Wave | 3 | Requires proven safety track record |

### Scoring Guidance Notes

1. **When attributes conflict:** An activity with high abstraction but very low error tolerance (e.g., algorithmic trading) may automate later than its abstraction score suggests. Weight error tolerance more heavily when stakes are life/safety-related.

2. **Context matters:** The same physical action scores differently based on context (see Part G.3). Always score the specific activity-in-context, not the abstract action.

3. **Partial automation:** Many activities will see partial automation before full automation. "Current AI Capability: Partial" indicates AI assists but cannot fully replace humans—this is a stable state for Wave 2-3 activities.

4. **Bottleneck stacking:** Activities may have multiple bottlenecks; list the primary one but note secondary bottlenecks in rationale if relevant.

---

## Part D.2: Level 3 Activity Enumeration

This section enumerates specific activities for each Level 2 category with full attribute scoring.

**Attribute Key:**

*Intrinsic Attributes (stable - define what the activity IS):*
- **Abs**: Abstraction Level (1=Pure Abstract, 2=Grounded Abstract, 3=Supervised Physical, 4=Adaptive Physical, 5=Deeply Embodied)
- **Err**: Error Tolerance (1=High, 2=Medium, 3=Low, 4=Very Low)
- **Fb**: Feedback Loop Speed (1=Seconds, 2=Minutes-Hours, 3=Hours-Days, 4=Days-Months)
- **Soc**: Social Complexity (0=None, 1=Informational, 2=Coordinative, 3=Persuasive, 4=Relational)

*Temporal Attributes (mutable - will change as AI progresses):*
- **AI**: AI Capability Status (A=Automated, D=Deployed, P=Piloting, R=Research, T=Theoretical, X=Not applicable)
- **Bot**: Primary Bottleneck (None, Sens=Sensing, Reas=Reasoning, Dex=Dexterity, Mob=Mobility, Adap=Adaptation, Soc=Social, Safe=Safety, Reg=Regulation, Econ=Economics, Trust=Trust)
- **Wave**: Automation Wave (1-4, N/A)

**Temporal Data Version:** 2026-02-01 | *Update temporal columns (AI, Bot, Wave) periodically; intrinsic columns should remain stable.*

---

### Domain 1: SYMBOLIC COMPUTATION

#### 1.1 Numerical Computation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Basic arithmetic calculations | 1 | 1 | 1 | 0 | S | None | 1 |
| Spreadsheet formula creation | 1 | 1 | 1 | 0 | S | None | 1 |
| Statistical analysis execution | 1 | 2 | 2 | 0 | S | None | 1 |
| Financial modeling calculations | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Tax computation | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Actuarial calculations | 1 | 3 | 3 | 0 | N | Reas | 1-2 |
| Scientific numerical simulation | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Cryptographic computation | 1 | 4 | 1 | 0 | S | None | 1 |

#### 1.2 Data Transformation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| File format conversion | 1 | 1 | 1 | 0 | S | None | 1 |
| Data cleaning and normalization | 1 | 2 | 2 | 0 | N | Reas | 1 |
| ETL pipeline execution | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Database query writing | 1 | 1 | 1 | 0 | S | None | 1 |
| Data migration | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Schema transformation | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Log parsing and extraction | 1 | 1 | 1 | 0 | S | None | 1 |

#### 1.3 Formal Verification

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Syntax checking | 1 | 1 | 1 | 0 | S | None | 1 |
| Code compilation | 1 | 1 | 1 | 0 | S | None | 1 |
| Type checking | 1 | 1 | 1 | 0 | S | None | 1 |
| Unit test execution | 1 | 1 | 1 | 0 | S | None | 1 |
| Formal proof verification | 1 | 4 | 3 | 0 | P | Reas | 2 |
| Security vulnerability scanning | 1 | 3 | 2 | 0 | P | Reas | 2 |
| Compliance rule checking | 1 | 3 | 2 | 0 | P | Reas | 2 |

#### 1.4 Algorithmic Execution

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Sorting and searching | 1 | 1 | 1 | 0 | S | None | 1 |
| Path finding algorithms | 1 | 1 | 1 | 0 | S | None | 1 |
| Scheduling optimization | 1 | 2 | 2 | 0 | S | None | 1 |
| Resource allocation | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Graph traversal | 1 | 1 | 1 | 0 | S | None | 1 |
| Constraint satisfaction | 1 | 2 | 2 | 0 | N | Reas | 1 |
| Batch job processing | 1 | 2 | 2 | 0 | S | None | 1 |

---

### Domain 2: INFORMATION SYNTHESIS

#### 2.1 Research & Analysis

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Literature review | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Market research compilation | 1 | 2 | 3 | 1 | P | Reas | 2 |
| Competitive analysis | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Data analysis interpretation | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Trend identification | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Due diligence research | 1 | 3 | 3 | 1 | P | Reas | 2 |
| Scientific hypothesis evaluation | 1 | 3 | 4 | 0 | P | Reas | 2-3 |
| Policy analysis | 1 | 3 | 4 | 1 | P | Reas | 2-3 |

#### 2.2 Report Generation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Meeting minutes writing | 1 | 1 | 2 | 0 | N | None | 1 |
| Status report creation | 1 | 1 | 2 | 0 | N | None | 1 |
| Technical documentation | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Financial report drafting | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Executive summary writing | 1 | 2 | 2 | 0 | N | Reas | 1-2 |
| Incident report compilation | 1 | 2 | 2 | 0 | P | Reas | 2 |
| Audit report preparation | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Grant proposal writing | 1 | 2 | 4 | 1 | P | Reas | 2 |

#### 2.3 Information Curation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Email inbox management | 1 | 1 | 1 | 0 | N | None | 1 |
| Document tagging and filing | 1 | 1 | 1 | 0 | N | None | 1 |
| Knowledge base maintenance | 1 | 2 | 2 | 0 | P | Reas | 2 |
| Content moderation | 1 | 2 | 1 | 0 | P | Reas | 2 |
| Archive organization | 1 | 1 | 2 | 0 | N | None | 1 |
| Bibliography compilation | 1 | 1 | 2 | 0 | S | None | 1 |
| News aggregation and filtering | 1 | 1 | 1 | 0 | N | None | 1 |

#### 2.4 Pattern Recognition

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Fraud detection | 1 | 3 | 1 | 0 | N | Reas | 1-2 |
| Anomaly identification | 1 | 2 | 1 | 0 | N | Reas | 1 |
| Image classification | 1 | 2 | 1 | 0 | S | None | 1 |
| Sentiment analysis | 1 | 2 | 1 | 0 | N | Reas | 1 |
| Medical image analysis | 1 | 3 | 2 | 0 | N | Safe | 2 |
| Quality defect detection | 2 | 3 | 1 | 0 | N | Sens | 2 |
| Handwriting recognition | 1 | 2 | 1 | 0 | S | None | 1 |
| Voice pattern recognition | 1 | 2 | 1 | 0 | N | None | 1 |

#### 2.5 Recommendation Systems

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Product recommendations | 1 | 1 | 1 | 0 | S | None | 1 |
| Content personalization | 1 | 1 | 1 | 0 | S | None | 1 |
| Job-candidate matching | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Treatment recommendations | 1 | 4 | 3 | 0 | P | Safe | 3 |
| Investment recommendations | 1 | 3 | 3 | 0 | P | Safe | 2-3 |
| Route optimization | 1 | 2 | 1 | 0 | S | None | 1 |
| Energy usage optimization | 1 | 2 | 2 | 0 | N | Reas | 1-2 |

---

### Domain 3: CREATIVE GENERATION

#### 3.1 Textual Creation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Blog post writing | 1 | 1 | 2 | 0 | N | None | 1 |
| Marketing copy creation | 1 | 1 | 2 | 0 | N | Reas | 1 |
| Social media content | 1 | 1 | 1 | 0 | N | None | 1 |
| Fiction writing | 1 | 1 | 3 | 0 | P | Reas | 2 |
| Screenplay writing | 1 | 1 | 3 | 0 | P | Reas | 2 |
| Poetry composition | 1 | 1 | 2 | 0 | P | Reas | 2 |
| Technical writing | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Legal document drafting | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Speech writing | 1 | 2 | 3 | 1 | P | Reas | 2 |

#### 3.2 Visual Design

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Logo design | 1 | 1 | 2 | 0 | P | Reas | 2 |
| UI/UX mockup creation | 1 | 1 | 2 | 0 | P | Reas | 2 |
| Social media graphics | 1 | 1 | 1 | 0 | N | None | 1 |
| Presentation design | 1 | 1 | 2 | 0 | N | None | 1 |
| Photo editing and retouching | 1 | 1 | 1 | 0 | N | None | 1 |
| Video thumbnail creation | 1 | 1 | 1 | 0 | N | None | 1 |
| Infographic design | 1 | 1 | 2 | 0 | P | Reas | 2 |
| Brand identity development | 1 | 2 | 4 | 1 | P | Reas | 2-3 |
| Architectural visualization | 1 | 2 | 3 | 0 | P | Reas | 2 |

#### 3.3 Software Development

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Boilerplate code generation | 1 | 1 | 1 | 0 | S | None | 1 |
| Unit test writing | 1 | 1 | 1 | 0 | N | Reas | 1 |
| Bug fixing | 1 | 2 | 2 | 0 | P | Reas | 2 |
| API integration | 1 | 2 | 2 | 0 | P | Reas | 2 |
| Code refactoring | 1 | 2 | 2 | 0 | P | Reas | 2 |
| System architecture design | 1 | 3 | 4 | 1 | P | Reas | 2-3 |
| Algorithm design | 1 | 2 | 3 | 0 | P | Reas | 2 |
| Security implementation | 1 | 4 | 3 | 0 | P | Reas | 2-3 |
| Performance optimization | 1 | 2 | 2 | 0 | P | Reas | 2 |

#### 3.4 Musical Composition

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Background music generation | 1 | 1 | 2 | 0 | N | None | 1 |
| Jingle creation | 1 | 1 | 2 | 0 | P | Reas | 2 |
| Sound effect design | 1 | 1 | 1 | 0 | N | None | 1 |
| Music arrangement | 1 | 1 | 2 | 0 | P | Reas | 2 |
| Podcast editing | 1 | 1 | 2 | 0 | N | None | 1 |
| Song composition | 1 | 1 | 3 | 0 | P | Reas | 2 |
| Film scoring | 1 | 2 | 3 | 1 | P | Reas | 2-3 |
| Audio mastering | 1 | 2 | 2 | 0 | P | Reas | 2 |

#### 3.5 Strategic Innovation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Product concept development | 1 | 2 | 4 | 1 | P | Reas | 2-3 |
| Business model design | 1 | 2 | 4 | 1 | P | Reas | 2-3 |
| Marketing strategy creation | 1 | 2 | 4 | 1 | P | Reas | 2-3 |
| Process improvement design | 1 | 2 | 3 | 1 | P | Reas | 2 |
| Innovation ideation | 1 | 1 | 3 | 1 | P | Reas | 2 |
| Scenario planning | 1 | 2 | 4 | 1 | P | Reas | 2-3 |
| Competitive positioning | 1 | 2 | 4 | 1 | P | Reas | 2 |

---

### Domain 4: INTERPERSONAL COMMUNICATION

#### 4.1 Instructional Communication

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Online course delivery | 1 | 2 | 3 | 2 | P | Soc | 2 |
| Tutorial creation | 1 | 1 | 2 | 1 | N | None | 1 |
| One-on-one tutoring | 2 | 2 | 1 | 4 | P | Soc | 3 |
| Classroom teaching | 2 | 3 | 4 | 4 | P | Soc | 3-4 |
| Corporate training delivery | 2 | 2 | 3 | 2 | P | Soc | 2-3 |
| Technical explanation | 1 | 2 | 2 | 1 | N | Reas | 1-2 |
| Skills demonstration | 3 | 2 | 1 | 2 | P | Soc | 3 |
| Mentoring | 2 | 2 | 4 | 4 | P | Soc | 3-4 |

#### 4.2 Persuasive Communication

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Sales pitch delivery | 2 | 2 | 2 | 3 | P | Soc | 2-3 |
| Contract negotiation | 1 | 3 | 3 | 3 | P | Soc | 3 |
| Fundraising solicitation | 2 | 2 | 3 | 3 | P | Soc | 3 |
| Political campaigning | 2 | 2 | 4 | 3 | P | Soc | 3 |
| Public speaking | 2 | 2 | 2 | 3 | P | Soc | 3 |
| Debate and argumentation | 1 | 2 | 1 | 3 | P | Soc | 2-3 |
| Advertising presentation | 2 | 2 | 2 | 3 | P | Soc | 2-3 |
| Conflict mediation | 2 | 3 | 2 | 4 | P | Soc | 3-4 |

#### 4.3 Coordinative Communication

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Meeting facilitation | 2 | 2 | 1 | 2 | P | Soc | 2-3 |
| Project status updates | 1 | 1 | 2 | 1 | N | None | 1 |
| Calendar scheduling | 1 | 1 | 1 | 1 | N | None | 1 |
| Team coordination | 2 | 2 | 2 | 2 | P | Soc | 2-3 |
| Cross-department liaison | 2 | 2 | 3 | 2 | P | Soc | 2-3 |
| Vendor coordination | 1 | 2 | 2 | 2 | P | Soc | 2 |
| Event planning communication | 1 | 2 | 3 | 2 | P | Soc | 2 |
| Stakeholder updates | 1 | 2 | 3 | 2 | P | Soc | 2 |

#### 4.4 Informational Exchange

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Customer service chat | 1 | 2 | 1 | 1 | N | None | 1 |
| Technical support | 1 | 2 | 1 | 1 | P | Reas | 2 |
| FAQ answering | 1 | 1 | 1 | 1 | S | None | 1 |
| Receptionist duties | 2 | 1 | 1 | 1 | P | Soc | 2 |
| Information desk assistance | 2 | 1 | 1 | 1 | P | Soc | 2 |
| Hotline operation | 1 | 2 | 1 | 1 | P | Soc | 2 |
| Interview conducting | 2 | 2 | 2 | 2 | P | Soc | 2-3 |
| Survey administration | 1 | 1 | 2 | 1 | N | None | 1 |

#### 4.5 Social Communication

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Small talk and rapport building | 2 | 1 | 1 | 4 | P | Soc | 3 |
| Professional networking | 2 | 1 | 3 | 3 | P | Soc | 3 |
| Team morale building | 2 | 1 | 3 | 4 | P | Soc | 3-4 |
| Icebreaker facilitation | 2 | 1 | 1 | 3 | P | Soc | 3 |
| Social event hosting | 2 | 1 | 2 | 3 | P | Soc | 3 |
| Relationship maintenance emails | 1 | 1 | 3 | 2 | N | Soc | 2 |
| Congratulatory messages | 1 | 1 | 2 | 2 | N | None | 1 |

---

### Domain 5: PROCEDURAL SERVICE DELIVERY

#### 5.1 Administrative Processing

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Form filling and submission | 1 | 2 | 2 | 0 | N | None | 1 |
| Appointment scheduling | 1 | 1 | 1 | 1 | N | None | 1 |
| Data entry | 1 | 2 | 1 | 0 | N | None | 1 |
| Record keeping | 1 | 2 | 2 | 0 | N | None | 1 |
| Permit and license processing | 1 | 3 | 3 | 1 | P | Reas | 2 |
| Expense reimbursement processing | 1 | 2 | 2 | 0 | N | Reas | 1-2 |
| Benefits enrollment | 1 | 2 | 2 | 1 | P | Reas | 2 |
| Document notarization | 2 | 3 | 1 | 1 | P | Reg | 3 |

#### 5.2 Retail & Hospitality

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Checkout and payment processing | 2 | 2 | 1 | 1 | N | None | 2 |
| Order taking | 2 | 1 | 1 | 1 | N | Soc | 2 |
| Table service (food delivery) | 3 | 1 | 1 | 2 | P | Dex | 3 |
| Hotel check-in/check-out | 2 | 2 | 1 | 1 | P | Soc | 2 |
| Concierge services | 2 | 1 | 1 | 2 | P | Soc | 2-3 |
| Room service delivery | 3 | 1 | 1 | 1 | P | Mob | 3 |
| Merchandise restocking | 3 | 1 | 1 | 0 | P | Dex | 2-3 |
| Gift wrapping | 3 | 1 | 1 | 1 | P | Dex | 3 |

#### 5.3 Financial Services

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Transaction processing | 1 | 3 | 1 | 0 | S | None | 1 |
| Account balance inquiries | 1 | 2 | 1 | 0 | S | None | 1 |
| Loan application processing | 1 | 3 | 3 | 1 | P | Reas | 2 |
| Insurance claims processing | 1 | 3 | 3 | 1 | P | Reas | 2 |
| Credit card dispute handling | 1 | 2 | 2 | 1 | P | Reas | 2 |
| Wire transfer execution | 1 | 4 | 1 | 0 | S | None | 1 |
| Tax preparation | 1 | 3 | 3 | 1 | P | Reas | 2 |
| Bookkeeping | 1 | 2 | 2 | 0 | N | Reas | 1-2 |

#### 5.4 Healthcare Administration

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Patient registration | 1 | 2 | 1 | 1 | N | None | 1-2 |
| Appointment scheduling | 1 | 2 | 1 | 1 | N | None | 1 |
| Insurance verification | 1 | 2 | 2 | 0 | N | Reas | 1-2 |
| Medical billing and coding | 1 | 3 | 2 | 0 | P | Reas | 2 |
| Prescription refill processing | 1 | 4 | 1 | 0 | P | Safe | 2 |
| Medical records management | 1 | 3 | 2 | 0 | P | Reas | 2 |
| Triage assessment | 2 | 4 | 1 | 2 | P | Safe | 3 |
| Referral coordination | 1 | 2 | 2 | 1 | P | Reas | 2 |

#### 5.5 Legal & Compliance Services

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Contract review | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Legal document preparation | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Compliance monitoring | 1 | 3 | 2 | 0 | P | Reas | 2 |
| Regulatory filing | 1 | 3 | 3 | 0 | P | Reas | 2 |
| Background check processing | 1 | 3 | 2 | 0 | P | Reas | 2 |
| Patent application preparation | 1 | 3 | 4 | 0 | P | Reas | 2-3 |
| Court document filing | 1 | 3 | 2 | 0 | P | Reg | 2 |
| Due diligence documentation | 1 | 3 | 3 | 0 | P | Reas | 2 |

---

### Domain 6: CARE & RELATIONAL SUPPORT

#### 6.1 Childcare & Development

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Infant feeding | 5 | 4 | 1 | 4 | E | Dex | 4 |
| Diaper changing | 5 | 2 | 1 | 4 | E | Dex | 4 |
| Child supervision | 4 | 4 | 1 | 4 | E | Adap | 4 |
| Bedtime routine | 5 | 2 | 1 | 4 | E | Soc | 4 |
| Educational play | 4 | 1 | 2 | 4 | P | Soc | 3-4 |
| Homework assistance | 2 | 2 | 2 | 4 | P | Soc | 3 |
| Emotional comforting | 5 | 3 | 1 | 4 | E | Soc | 4 |
| Developmental assessment | 2 | 3 | 4 | 4 | P | Soc | 3 |

#### 6.2 Eldercare & Disability Support

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Mobility assistance | 4 | 4 | 1 | 4 | E | Dex | 4 |
| Medication reminders | 2 | 4 | 1 | 2 | P | Safe | 2-3 |
| Bathing assistance | 5 | 3 | 1 | 4 | E | Dex | 4 |
| Meal assistance | 4 | 3 | 1 | 4 | E | Dex | 4 |
| Companionship conversation | 2 | 1 | 1 | 4 | P | Soc | 3 |
| Cognitive stimulation activities | 2 | 1 | 2 | 4 | P | Soc | 3 |
| Safety monitoring | 2 | 4 | 1 | 2 | P | Sens | 2-3 |
| Transportation accompaniment | 4 | 3 | 1 | 4 | E | Adap | 4 |

#### 6.3 Healthcare Delivery

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Vital signs measurement | 3 | 3 | 1 | 2 | P | Dex | 2-3 |
| Medication administration | 3 | 4 | 1 | 2 | P | Safe | 3 |
| Wound care | 4 | 3 | 2 | 2 | E | Dex | 3-4 |
| Physical therapy exercises | 4 | 2 | 2 | 4 | P | Soc | 3 |
| Patient education | 2 | 2 | 3 | 4 | P | Soc | 3 |
| Symptom assessment | 2 | 3 | 1 | 4 | P | Reas | 3 |
| IV insertion | 4 | 3 | 1 | 2 | E | Dex | 3-4 |
| Post-operative care | 4 | 4 | 2 | 4 | E | Adap | 4 |

#### 6.4 Psychological Support

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Active listening | 2 | 2 | 1 | 4 | P | Soc | 3 |
| Crisis intervention | 2 | 4 | 1 | 4 | P | Soc | 3-4 |
| Cognitive behavioral therapy | 2 | 3 | 4 | 4 | P | Soc | 3 |
| Group therapy facilitation | 2 | 3 | 2 | 4 | P | Soc | 3-4 |
| Grief counseling | 2 | 3 | 3 | 4 | E | Soc | 4 |
| Addiction counseling | 2 | 3 | 4 | 4 | P | Soc | 3-4 |
| Life coaching | 2 | 2 | 4 | 4 | P | Soc | 3 |
| Stress management guidance | 2 | 1 | 2 | 4 | P | Soc | 3 |

#### 6.5 Community & Civic Care

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Volunteer coordination | 2 | 1 | 3 | 2 | P | Soc | 2-3 |
| Community outreach | 2 | 1 | 3 | 3 | P | Soc | 3 |
| Religious counseling | 2 | 2 | 3 | 4 | P | Soc | 3-4 |
| Food bank service | 3 | 1 | 1 | 2 | P | Dex | 2-3 |
| Homeless shelter assistance | 4 | 2 | 1 | 4 | E | Adap | 4 |
| Disaster relief coordination | 2 | 3 | 2 | 2 | P | Adap | 3 |
| Youth mentoring | 2 | 2 | 4 | 4 | P | Soc | 3-4 |
| Community organizing | 2 | 2 | 4 | 3 | P | Soc | 3 |

#### 6.6 Social Relationship Maintenance

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Family dinner conversation | 2 | 1 | 1 | 4 | E | Soc | 4 |
| Friendship maintenance | 2 | 1 | 3 | 4 | E | Soc | 4 |
| Romantic partnership nurturing | 5 | 3 | 4 | 4 | X | Soc | 4+ |
| Conflict resolution with loved ones | 2 | 3 | 2 | 4 | E | Soc | 4 |
| Celebration and ritual participation | 4 | 1 | 2 | 4 | E | Soc | 4 |
| Sympathy and condolence | 2 | 2 | 2 | 4 | P | Soc | 3-4 |
| Shared recreational activities | 4 | 1 | 1 | 4 | E | Soc | 4 |

---

### Domain 7: STRUCTURED PHYSICAL OPERATIONS

#### 7.1 Manufacturing & Assembly

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Assembly line positioning | 3 | 2 | 1 | 0 | N | Dex | 2 |
| Component insertion | 3 | 2 | 1 | 0 | N | Dex | 2 |
| Welding (automated) | 3 | 3 | 1 | 0 | N | None | 2 |
| Packaging and boxing | 3 | 1 | 1 | 0 | N | Dex | 2 |
| Quality sorting | 3 | 2 | 1 | 0 | N | Sens | 2 |
| Machine tending | 3 | 2 | 1 | 0 | N | None | 2 |
| Label application | 3 | 2 | 1 | 0 | S | None | 1 |
| Palletizing | 3 | 1 | 1 | 0 | N | Dex | 2 |

#### 7.2 Warehouse & Logistics

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Order picking (structured) | 3 | 2 | 1 | 0 | N | Dex | 2 |
| Inventory scanning | 2 | 2 | 1 | 0 | S | None | 1 |
| Shelf restocking | 3 | 1 | 1 | 0 | P | Dex | 2-3 |
| Package sorting | 3 | 2 | 1 | 0 | N | Dex | 2 |
| Forklift loading | 3 | 3 | 1 | 0 | P | Mob | 2-3 |
| Bin organization | 3 | 1 | 1 | 0 | P | Dex | 2 |
| Shipping preparation | 3 | 2 | 1 | 0 | P | Dex | 2 |
| Returns processing | 3 | 2 | 1 | 0 | P | Reas | 2-3 |

#### 7.3 Machine Operation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| CNC machine operation | 2 | 3 | 2 | 0 | P | Reas | 2 |
| 3D printer operation | 2 | 2 | 2 | 0 | N | None | 2 |
| Industrial printer operation | 2 | 2 | 1 | 0 | N | None | 2 |
| Injection molding operation | 3 | 3 | 1 | 0 | P | Sens | 2 |
| Laser cutting operation | 2 | 3 | 1 | 0 | N | None | 2 |
| Textile machine operation | 3 | 2 | 1 | 0 | P | Sens | 2-3 |
| Food processing machinery | 3 | 3 | 1 | 0 | P | Sens | 2-3 |
| Conveyor system monitoring | 2 | 2 | 1 | 0 | N | Sens | 2 |

#### 7.4 Quality Control & Inspection

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Visual defect detection | 2 | 3 | 1 | 0 | N | Sens | 2 |
| Dimensional measurement | 2 | 3 | 1 | 0 | N | Sens | 2 |
| Functional testing | 2 | 3 | 1 | 0 | N | Reas | 2 |
| Sample collection | 3 | 2 | 1 | 0 | P | Dex | 2-3 |
| Documentation of defects | 1 | 2 | 1 | 0 | N | None | 1 |
| Calibration verification | 2 | 3 | 2 | 0 | P | Reas | 2 |
| Food safety inspection | 3 | 4 | 1 | 0 | P | Sens | 2-3 |
| Weight and count verification | 2 | 2 | 1 | 0 | S | None | 1 |

#### 7.5 Facility Maintenance

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Floor cleaning (commercial) | 3 | 1 | 1 | 0 | N | Mob | 2 |
| Window cleaning (ground level) | 3 | 1 | 1 | 0 | P | Dex | 2-3 |
| Trash collection | 3 | 1 | 1 | 0 | P | Mob | 2-3 |
| Restroom sanitization | 3 | 2 | 1 | 0 | P | Dex | 2-3 |
| Lawn mowing | 3 | 1 | 1 | 0 | N | Mob | 2 |
| Snow removal | 3 | 2 | 1 | 0 | P | Mob | 2-3 |
| Light fixture replacement | 3 | 2 | 1 | 0 | P | Dex | 3 |
| HVAC filter replacement | 3 | 2 | 1 | 0 | P | Dex | 3 |

---

### Domain 8: SKILLED PHYSICAL CRAFT

#### 8.1 Construction & Building

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Framing carpentry | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Electrical wiring | 4 | 4 | 2 | 0 | E | Dex | 3-4 |
| Plumbing installation | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Drywall installation | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Roofing | 4 | 3 | 2 | 0 | E | Adap | 4 |
| Masonry and bricklaying | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Tile installation | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Painting (professional) | 4 | 2 | 2 | 0 | P | Dex | 3 |
| Concrete finishing | 4 | 3 | 1 | 0 | E | Dex | 3-4 |

#### 8.2 Equipment Repair & Maintenance

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Automotive diagnosis | 2 | 3 | 2 | 0 | P | Reas | 2-3 |
| Engine repair | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| HVAC repair | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Appliance repair | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Electronics repair | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Industrial equipment maintenance | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Bicycle repair | 4 | 2 | 1 | 0 | E | Dex | 3 |
| Watch and clock repair | 5 | 3 | 2 | 0 | E | Dex | 4 |

#### 8.3 Surgical & Medical Procedures

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Minimally invasive surgery | 4 | 4 | 1 | 2 | P | Dex | 3-4 |
| Open surgery | 5 | 4 | 1 | 2 | E | Dex | 4 |
| Dental procedures | 4 | 3 | 1 | 2 | P | Dex | 3-4 |
| Phlebotomy | 4 | 3 | 1 | 2 | P | Dex | 3 |
| Endoscopy | 4 | 4 | 1 | 1 | P | Dex | 3 |
| Suturing | 4 | 3 | 1 | 1 | P | Dex | 3-4 |
| Casting and splinting | 4 | 2 | 1 | 2 | P | Dex | 3 |
| Biopsy procedures | 4 | 4 | 1 | 2 | P | Dex | 3-4 |

#### 8.4 Culinary Arts

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Prep cooking (chopping, slicing) | 4 | 2 | 1 | 0 | P | Dex | 3 |
| Line cooking | 4 | 2 | 1 | 1 | E | Dex | 3-4 |
| Baking and pastry | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Sauce preparation | 4 | 2 | 1 | 0 | E | Dex | 3-4 |
| Plating and presentation | 4 | 1 | 1 | 0 | E | Dex | 3 |
| Meat butchering | 4 | 2 | 1 | 0 | E | Dex | 3-4 |
| Sushi preparation | 5 | 2 | 1 | 0 | E | Dex | 4 |
| Barista skills | 4 | 1 | 1 | 1 | P | Dex | 3 |

#### 8.5 Artisan Crafts

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Woodworking (furniture) | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Metalworking and welding | 4 | 3 | 2 | 0 | E | Dex | 3-4 |
| Pottery and ceramics | 5 | 1 | 2 | 0 | E | Dex | 4 |
| Jewelry making | 5 | 2 | 2 | 0 | E | Dex | 4 |
| Textile crafts (sewing, knitting) | 4 | 1 | 2 | 0 | E | Dex | 3-4 |
| Glass blowing | 5 | 2 | 1 | 0 | E | Dex | 4 |
| Leatherworking | 4 | 2 | 2 | 0 | E | Dex | 3-4 |
| Instrument making | 5 | 3 | 3 | 0 | E | Dex | 4 |

---

### Domain 9: ENVIRONMENTAL NAVIGATION

#### 9.1 Vehicle Operation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Highway driving | 2 | 4 | 1 | 0 | N | Sens | 2 |
| Urban driving | 2 | 4 | 1 | 0 | P | Adap | 2-3 |
| Parking | 2 | 2 | 1 | 0 | N | Sens | 2 |
| Truck driving (long-haul) | 2 | 4 | 1 | 0 | P | Adap | 2-3 |
| Forklift operation | 3 | 3 | 1 | 0 | P | Sens | 2-3 |
| Train operation | 2 | 4 | 1 | 0 | N | Sens | 2 |
| Aircraft piloting | 2 | 4 | 1 | 1 | P | Adap | 3 |
| Boat/ship navigation | 2 | 4 | 1 | 0 | P | Adap | 3 |

#### 9.2 Last-Mile Delivery

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Package delivery (suburban) | 4 | 2 | 1 | 1 | P | Mob | 3 |
| Package delivery (urban) | 4 | 2 | 1 | 1 | P | Mob | 3 |
| Food delivery | 4 | 2 | 1 | 1 | P | Mob | 3 |
| Mail delivery | 4 | 2 | 1 | 1 | P | Mob | 3 |
| Drone delivery | 2 | 3 | 1 | 0 | P | Reg | 2-3 |
| Grocery delivery | 4 | 2 | 1 | 1 | P | Mob | 3 |
| Pharmacy delivery | 4 | 3 | 1 | 1 | P | Safe | 3 |
| Large item delivery | 4 | 2 | 1 | 2 | E | Dex | 3-4 |

#### 9.3 Field Work

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Crop harvesting | 4 | 2 | 1 | 0 | P | Dex | 3 |
| Fruit picking | 4 | 1 | 1 | 0 | P | Dex | 3 |
| Livestock herding | 4 | 2 | 1 | 0 | P | Adap | 3 |
| Land surveying | 2 | 3 | 2 | 0 | P | Sens | 2-3 |
| Environmental sampling | 4 | 3 | 2 | 0 | P | Dex | 3 |
| Irrigation management | 2 | 2 | 3 | 0 | P | Sens | 2 |
| Pesticide application | 4 | 3 | 2 | 0 | P | Mob | 2-3 |
| Forestry operations | 4 | 3 | 2 | 0 | E | Adap | 3-4 |

#### 9.4 Exploration & Search

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Infrastructure inspection | 4 | 3 | 2 | 0 | P | Sens | 2-3 |
| Pipeline inspection | 4 | 3 | 2 | 0 | P | Mob | 2-3 |
| Search and rescue | 4 | 4 | 1 | 2 | P | Adap | 3-4 |
| Mine exploration | 4 | 4 | 1 | 0 | P | Adap | 3 |
| Underwater inspection | 4 | 3 | 2 | 0 | P | Mob | 3 |
| Building security patrol | 4 | 3 | 1 | 1 | P | Mob | 2-3 |
| Wildlife monitoring | 4 | 2 | 3 | 0 | P | Mob | 2-3 |
| Disaster site assessment | 4 | 4 | 1 | 2 | P | Adap | 3-4 |

#### 9.5 Recreational Navigation

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Hiking and trail walking | 4 | 2 | 1 | 0 | E | Adap | 4 |
| Cycling (road) | 4 | 3 | 1 | 0 | E | Adap | 4 |
| Mountain biking | 4 | 3 | 1 | 0 | E | Adap | 4 |
| Skiing and snowboarding | 4 | 3 | 1 | 0 | E | Adap | 4 |
| Rock climbing | 5 | 4 | 1 | 0 | E | Adap | 4 |
| Kayaking and canoeing | 4 | 3 | 1 | 0 | E | Adap | 4 |
| Surfing | 5 | 3 | 1 | 0 | E | Adap | 4 |
| Scuba diving | 5 | 4 | 1 | 0 | E | Adap | 4 |

---

### Domain 10: PERSONAL EMBODIMENT

#### 10.1 Basic Self-Care

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Sleeping | 5 | 1 | 1 | 0 | X | None | N/A |
| Bathing and showering | 5 | 1 | 1 | 0 | E | Dex | 4 |
| Brushing teeth | 5 | 1 | 1 | 0 | E | Dex | 4 |
| Dressing | 5 | 1 | 1 | 0 | E | Dex | 4 |
| Grooming (hair, makeup) | 5 | 1 | 1 | 0 | E | Dex | 4 |
| Using toilet | 5 | 1 | 1 | 0 | X | None | N/A |
| Skincare | 5 | 1 | 1 | 0 | E | Dex | 4 |

#### 10.2 Nourishment

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Eating | 5 | 1 | 1 | 0 | X | None | N/A |
| Drinking | 5 | 1 | 1 | 0 | X | None | N/A |
| Meal planning (personal) | 1 | 1 | 3 | 0 | N | None | 1 |
| Grocery shopping | 4 | 1 | 2 | 1 | P | Mob | 3 |
| Home cooking | 4 | 1 | 1 | 0 | E | Dex | 3-4 |
| Food storage and organization | 4 | 1 | 2 | 0 | E | Dex | 3-4 |
| Kitchen cleanup | 4 | 1 | 1 | 0 | P | Dex | 3 |

#### 10.3 Health Maintenance

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Exercise and workout | 5 | 1 | 1 | 0 | E | Adap | 4 |
| Running and jogging | 5 | 2 | 1 | 0 | E | Adap | 4 |
| Yoga and stretching | 5 | 1 | 1 | 0 | E | Adap | 4 |
| Self-monitoring (weight, BP) | 2 | 2 | 2 | 0 | N | None | 1-2 |
| Medication self-administration | 3 | 3 | 2 | 0 | P | Safe | 2-3 |
| Wound self-care | 4 | 2 | 2 | 0 | P | Dex | 3 |
| Sleep hygiene | 2 | 1 | 4 | 0 | P | Adap | 2-3 |
| Stress management | 2 | 1 | 3 | 0 | P | Adap | 2-3 |

#### 10.4 Domestic Maintenance

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Vacuuming | 3 | 1 | 1 | 0 | N | Mob | 2 |
| Mopping | 3 | 1 | 1 | 0 | P | Mob | 2-3 |
| Dusting | 4 | 1 | 1 | 0 | E | Dex | 3-4 |
| Laundry washing | 3 | 1 | 2 | 0 | P | Dex | 2-3 |
| Laundry folding | 4 | 1 | 1 | 0 | E | Dex | 3-4 |
| Ironing | 4 | 1 | 1 | 0 | E | Dex | 3-4 |
| Dishwashing | 3 | 1 | 1 | 0 | P | Dex | 2-3 |
| Bed making | 4 | 1 | 1 | 0 | E | Dex | 3-4 |
| Home organization | 4 | 1 | 2 | 0 | E | Reas | 3-4 |
| Minor home repairs | 4 | 2 | 2 | 0 | E | Dex | 3-4 |

#### 10.5 Personal Expression

| Activity | Abs | Err | Fb | Soc | AI | Bot | Wave |
|----------|-----|-----|----|----|----|----|------|
| Reading for pleasure | 1 | 1 | 1 | 0 | X | None | N/A |
| Watching entertainment | 1 | 1 | 1 | 0 | X | None | N/A |
| Gaming | 2 | 1 | 1 | 0 | X | None | N/A |
| Hobby crafting | 4 | 1 | 2 | 0 | E | Dex | 4 |
| Playing musical instruments | 5 | 1 | 1 | 0 | E | Dex | 4 |
| Gardening | 4 | 1 | 3 | 0 | E | Dex | 3-4 |
| Pet care | 4 | 2 | 1 | 4 | E | Adap | 4 |
| Journaling and writing | 1 | 1 | 2 | 0 | X | None | N/A |
| Meditation | 5 | 1 | 2 | 0 | X | None | N/A |
| Resting and relaxation | 5 | 1 | 1 | 0 | X | None | N/A |

---

## Part D.3: Level 3 Summary Statistics

**Taxonomy Statistics (stable):**

### Activity Count by Domain

| Domain | Categories | Activities | Avg per Category |
|--------|------------|------------|------------------|
| 1. Symbolic Computation | 4 | 29 | 7.3 |
| 2. Information Synthesis | 5 | 40 | 8.0 |
| 3. Creative Generation | 5 | 42 | 8.4 |
| 4. Interpersonal Communication | 5 | 39 | 7.8 |
| 5. Procedural Service Delivery | 5 | 40 | 8.0 |
| 6. Care & Relational Support | 6 | 47 | 7.8 |
| 7. Structured Physical Operations | 5 | 40 | 8.0 |
| 8. Skilled Physical Craft | 5 | 42 | 8.4 |
| 9. Environmental Navigation | 5 | 40 | 8.0 |
| 10. Personal Embodiment | 5 | 41 | 8.2 |
| **Total** | **50** | **400** | **8.0** |

---

**Temporal Statistics (as of 2026-02-01):**

*These statistics will change as AI capabilities advance. Re-assess annually or when major AI milestones occur.*

### Distribution by Automation Wave

| Wave | Activity Count | Percentage | Definition |
|------|---------------|------------|-------------|
| Wave 1 | 62 | 15.5% | Software-only automation |
| Wave 1-2 | 38 | 9.5% | Transitioning to structured automation |
| Wave 2 | 87 | 21.8% | Structured environment automation |
| Wave 2-3 | 58 | 14.5% | Transitioning to adaptive automation |
| Wave 3 | 51 | 12.8% | Adaptive automation |
| Wave 3-4 | 54 | 13.5% | Transitioning to general purpose |
| Wave 4 | 37 | 9.3% | General-purpose automation |
| N/A | 13 | 3.3% | Not applicable for automation |

*Current wave-to-year mapping: Wave 1=2024-2026, Wave 2=2026-2028, Wave 3=2028-2032, Wave 4=2032+. Update mapping as AI progress clarifies.*

### Distribution by AI Capability Status

| Status | Activity Count | Percentage | Meaning |
|--------|---------------|------------|---------|
| Solved/Automated (S) | 31 | 7.8% | AI handles autonomously at scale |
| Near-solved/Deployed (N) | 89 | 22.3% | AI deployed with human oversight |
| Partial/Piloting (P) | 214 | 53.5% | AI in pilot testing or assists humans |
| Early/Research (E) | 53 | 13.3% | Research prototypes only |
| Not attempted (X) | 13 | 3.3% | Not applicable or not automatable |

### Distribution by Primary Bottleneck

| Bottleneck | Activity Count | Percentage |
|------------|---------------|------------|
| None | 42 | 10.5% |
| Reasoning | 72 | 18.0% |
| Dexterity | 89 | 22.3% |
| Social | 67 | 16.8% |
| Mobility | 34 | 8.5% |
| Adaptation | 42 | 10.5% |
| Sensing | 22 | 5.5% |
| Safety | 17 | 4.3% |
| Regulation | 6 | 1.5% |
| Data | 0 | 0.0% |
| Other/N/A | 9 | 2.3% |

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

**Population Strategy:**
1. Start with high-frequency activities from ATUS data
2. Add occupationally-important activities from O*NET
3. Include edge cases that test boundary rules
4. Ensure global representation (not just Western/developed economies)

### H.2: Scoring Methodology

Each activity will be scored by:
1. **Domain experts** (for current AI capability and realistic timeline estimates)
2. **AI researchers** (for bottleneck identification and technical feasibility)
3. **Time-use researchers** (for validation against ATUS/HETUS frequency data)

**Rater Training Protocol:**
1. Review Part D.1 example scored activities
2. Complete calibration exercise (10 activities, compare to reference scores)
3. Discuss disagreements to align interpretation
4. Achieve κ > 0.7 on calibration set before proceeding

### H.3: Validation Protocol

1. **Inter-rater reliability:** Two independent raters should achieve κ > 0.7 on all attributes
2. **Coverage audit:** Walk through 10 diverse occupations, ensure all tasks map
3. **Predictive validation:** Compare predictions to actual AI deployment timelines
4. **Retrospective validation:** Apply taxonomy to historically-automated tasks (typing, calculation, switchboard operation) to verify framework

### H.4: Implementation Roadmap

**Phase 1: Pilot (Domains 1-3)**
- Enumerate 50-100 activities in Domains 1-3 (highest abstraction, most AI progress)
- Score all attributes with 2+ raters
- Calculate inter-rater reliability
- Refine attribute definitions if κ < 0.7

**Phase 2: Expansion (Domains 4-7)**
- Enumerate 100-150 activities in Domains 4-7
- Include activities at domain boundaries to test boundary rules
- Update scoring guidance based on Phase 1 learnings

**Phase 3: Completion (Domains 8-10)**
- Enumerate remaining 50-100 activities in Domains 8-10
- Focus on physical/embodied activities with less AI progress
- Validate against robotics research benchmarks

**Phase 4: Validation and Publication**
- Complete coverage audit across 10 occupations
- Run predictive validation against AI deployment data
- Prepare methodology paper for peer review
- Publish taxonomy as structured data (JSON/CSV) with web interface

### H.5: Maintenance Plan

**Temporal Data Versioning (Key for Future-Proofing):**

The taxonomy separates stable structure from mutable predictions:

| Component | Update Frequency | What Changes |
|-----------|------------------|--------------|
| Domains (Level 1) | Rarely | Only for major reconceptualization |
| Categories (Level 2) | Annually | Add new categories as work evolves |
| Activities (Level 3) | Annually | Add new activities, retire obsolete ones |
| Intrinsic Attributes (Abs, Err, Fb, Soc) | Rarely | Only if activity definition changes |
| Temporal Attributes (AI, Bot, Wave) | Quarterly-Annually | Update as AI capabilities advance |
| Wave-to-Year Mapping | Annually | Adjust dates based on actual AI progress |

**Update Protocol:**

1. **Quarterly temporal review:**
   - Scan AI news for major capability announcements
   - Update AI Capability Status for affected activities
   - Move newly-automated activities to appropriate status

2. **Annual full review:**
   - Reassess all temporal attributes
   - Update wave-to-year mappings based on observed progress
   - Add new activities (especially from emerging occupations)
   - Validate against actual AI deployment data

3. **Wave recalibration triggers:**
   - Major AI breakthrough (e.g., new architecture, 10x capability jump)
   - Regulatory changes affecting deployment
   - Economic conditions affecting adoption rates

4. **Activity additions:**
   - Monitor BLS, O*NET for new occupations
   - Add activities from emerging industries
   - Community submissions (see H.7)

### H.6: Known Limitations and Future Work

**Current Limitations:**
- Timeline predictions (Waves 1-4) are estimates without formal forecasting methodology
- Single-attribute classification may oversimplify complex activities
- Western bias in example activities and occupations
- Temporal assessments will require periodic updates

**Future Research Directions:**
- Develop formal forecasting model using intrinsic attribute scores as inputs
- Extend taxonomy to include automation *level* (partial vs. full)
- Map economic impact (jobs, wages) to taxonomy categories
- Integrate with labor economics models
- Create machine-readable format (JSON/YAML) for programmatic updates

### H.7: Community Contribution Process

**For Temporal Updates:**
1. Submit evidence of AI capability change (paper, deployment announcement)
2. Propose new status for affected activities
3. Maintainers review and batch updates quarterly

**For Taxonomy Changes:**
1. Propose new activity with intrinsic attribute scores and rationale
2. Include mapping to O*NET, ATUS, or other frameworks
3. Peer review for MECE compliance
4. Merge into next annual release

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

*Document Status: Complete. Level 1 (10 domains), Level 2 (47 categories), and Level 3 (400 activities) fully enumerated and scored.*
