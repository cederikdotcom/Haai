# HAAI - Human Activity Automation Index

A comprehensive taxonomy for predicting AI automation timelines across all human activities.

## Overview

HAAI classifies all human activities—paid work and unpaid life—into a structured taxonomy optimized for predicting when artificial intelligence will achieve human-level capability at each activity.

### Core Thesis

- **Software AGI** operates in abstract/symbolic space → solves abstract tasks first
- **Embodied AGI** requires physical interaction → solves physical tasks later
- **Primary predictors:** (1) Abstraction level, (2) Error tolerance

## File Structure

```
Haai/
├── taxonomy.json    # Domain and category definitions (10 domains, 50 categories)
├── scoring.json     # Attribute definitions and scoring scales
├── mappings.json    # External taxonomy mappings (O*NET, ATUS, ISCO, etc.)
├── validation.json  # Test cases and validation checklists
└── README.md        # This file
```

## Taxonomy Structure

### Domains (Level 1)

Ordered from **highest abstraction** (earliest AI automation) to **lowest abstraction** (latest AI automation):

| Domain | Abstraction | AGI Wave | Primary AI Type |
|--------|-------------|----------|-----------------|
| 1. Symbolic Computation | 10/10 | 1 | Rule engines, calculators |
| 2. Information Synthesis | 9/10 | 1-2 | Language models, RAG |
| 3. Creative Generation | 8/10 | 1-2 | Generative AI |
| 4. Interpersonal Communication | 7/10 | 2 | Conversational AI |
| 5. Procedural Service Delivery | 6/10 | 2-3 | Multimodal AI + workflows |
| 6. Care & Relational Support | 5/10 | 3-4 | Affective computing |
| 7. Structured Physical Operations | 4/10 | 2-3 | Industrial robotics |
| 8. Skilled Physical Craft | 3/10 | 3-4 | Dexterous manipulation |
| 9. Environmental Navigation | 2/10 | 2-4 | Mobile robotics, AV |
| 10. Personal Embodiment | 1/10 | 3-4 | General-purpose robots |

### Scoring Attributes

Each activity is scored on seven dimensions:

1. **Abstraction Level** (1-5): Pure abstract → Deeply embodied
2. **Error Tolerance** (1-4): High tolerance → Very low tolerance
3. **Feedback Loop Speed** (1-4): Immediate → Long-term
4. **Social Complexity** (0-4): None → Relational
5. **Current AI Capability**: solved | near_solved | partial | early | not_attempted
6. **Primary Bottleneck**: sensing | reasoning | dexterity | mobility | adaptation | social | safety | regulation | data
7. **Estimated AGI Wave** (1-4): 2024-2026 → 2032+

## Usage

### Loading the Taxonomy

```javascript
const taxonomy = require('./taxonomy.json');
const scoring = require('./scoring.json');

// Get all domains
taxonomy.domains.forEach(domain => {
  console.log(`${domain.id}. ${domain.name} (Abstraction: ${domain.abstractionScore}/10)`);
});

// Get scoring scales
scoring.attributes.forEach(attr => {
  console.log(`${attr.name}: ${attr.description}`);
});
```

```python
import json

with open('taxonomy.json') as f:
    taxonomy = json.load(f)

with open('scoring.json') as f:
    scoring = json.load(f)

# Access domains
for domain in taxonomy['domains']:
    print(f"{domain['id']}. {domain['name']}")
```

### Classifying an Activity

1. Identify the **primary purpose** of the activity
2. Use **boundary rules** in `taxonomy.json` for adjacent domains
3. Apply **context modifiers** for professional vs personal activities
4. Score using scales in `scoring.json`

### Validation

Test cases in `validation.json` provide:
- Day-in-life coverage tests
- Occupation coverage tests
- AI capability alignment tests
- Edge case resolutions

## Design Principles

1. **Predict, don't describe** - Optimize for attributes that predict AI capability timing
2. **Task, not job** - Classify individual tasks, not occupations
3. **Observable, not inferred** - Activities are objectively identifiable
4. **Technology-agnostic** - Define by human capability required
5. **Globally applicable** - Avoid Western/developed-world bias
6. **MECE compliant** - Mutually exclusive, collectively exhaustive

## External Mappings

The taxonomy maps to:
- **O*NET** Work Activities
- **ATUS** (American Time Use Survey)
- **ISCO-08** Occupations
- **BEHAVIOR-1K** Robotic tasks
- **ActivityNet/Kinetics** Action classes

See `mappings.json` for detailed crosswalks.

## Roadmap

- [x] Level 1: Domain definitions (10 domains)
- [x] Level 2: Category definitions (50 categories)
- [ ] Level 3: Activity enumeration (200-400 activities with scores)
- [ ] Validation: Inter-rater reliability testing
- [ ] Predictive validation against AI deployment data

## Contributing

When adding activities:
1. Ensure MECE compliance (one category only)
2. Score all seven attributes
3. Add to validation test cases
4. Map to external taxonomies where applicable

## License

See LICENSE file.
