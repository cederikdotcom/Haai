# HAAI - Human Activity Automation Index

A structured taxonomy for predicting AI automation timelines across all human activities.

## Overview

HAAI classifies human activities—both paid work and unpaid life—into a hierarchical structure optimized for predicting when AI will achieve human-level capability at each activity.

**Core Thesis:**
- Software AGI operates in abstract/symbolic space → solves abstract tasks first
- Embodied AGI requires physical interaction → solves physical tasks later
- Primary predictors: **Abstraction level** and **Error tolerance**

## Structure

```
├── taxonomy.json     # Domains (10) and Categories (47) with classification rules
├── scoring.json      # Attribute definitions and scoring methodology
├── activities.json   # Level 3 activities with full attribute scores
├── mappings.json     # External system mappings (O*NET, ATUS, etc.)
└── Haai.md          # Original markdown documentation (reference)
```

## Domain Hierarchy

| Domain | Name | Abstraction | AGI Wave |
|--------|------|-------------|----------|
| 1 | Symbolic Computation | 10/10 | 1 |
| 2 | Information Synthesis | 9/10 | 1-2 |
| 3 | Creative Generation | 8/10 | 1-2 |
| 4 | Interpersonal Communication | 7/10 | 2 |
| 5 | Procedural Service Delivery | 6/10 | 2-3 |
| 6 | Care & Relational Support | 5/10 | 3-4 |
| 7 | Structured Physical Operations | 4/10 | 2-3 |
| 8 | Skilled Physical Craft | 3/10 | 3-4 |
| 9 | Environmental Navigation | 2/10 | 2-4 |
| 10 | Personal Embodiment | 1/10 | 3-4 |

## File Descriptions

### taxonomy.json
Core taxonomy structure containing:
- 10 domains with definitions, inclusion/exclusion criteria
- 47 categories organized under domains
- Classification rules for edge cases
- Boundary rules for domain disambiguation

### scoring.json
Attribute definitions for scoring activities:
- **Abstraction Level** (1-5): Pure abstract → Deeply embodied
- **Error Tolerance** (1-4): High → Very low
- **Feedback Loop Speed** (1-4): Immediate → Long-term
- **Social Complexity** (0-4): None → Relational
- **Current AI Capability**: Solved → Not attempted
- **Primary Bottleneck**: Sensing, Reasoning, Dexterity, etc.
- **Estimated AGI Wave** (1-4): 2024-2026 → 2032+

### activities.json
Level 3 activities with:
- Classification decision tree for resolving ambiguities
- 47 sample activities across all domains
- Full attribute scores for each activity
- Population guidelines for expansion

### mappings.json
Cross-references to external systems:
- O*NET Work Activities
- ATUS (American Time Use Survey)
- ISCO-08 occupational classifications
- BEHAVIOR-1K robotics benchmarks
- ActivityNet/Kinetics action recognition datasets

## Usage

### Classifying an Activity

1. Use the decision tree in `activities.json` to determine domain
2. Match to specific category based on domain criteria in `taxonomy.json`
3. Score attributes using scales defined in `scoring.json`
4. Validate against mappings to external systems

### Adding New Activities

```json
{
  "activityId": "3.3.3",
  "categoryId": "3.3",
  "name": "Activity name",
  "description": "What this activity involves",
  "abstractionLevel": 1,
  "errorTolerance": 2,
  "feedbackLoopSpeed": 1,
  "socialComplexity": 0,
  "currentAICapability": "partial",
  "primaryBottleneck": "reasoning",
  "estimatedAGIWave": 2
}
```

## Design Principles

1. **Predict, don't describe** - Optimize for AI capability timing
2. **Task, not job** - Decompose roles into constituent activities
3. **Observable, not inferred** - Objectively identifiable activities
4. **Technology-agnostic** - Define by human capability required
5. **Globally applicable** - Avoid cultural bias
6. **MECE compliant** - Every activity fits exactly one category

## Status

- **Level 1 (Domains):** Complete (10)
- **Level 2 (Categories):** Complete (47)
- **Level 3 (Activities):** Partial (47 of 200-400 target)

## References

- [O*NET Content Model](https://www.onetcenter.org/content.html)
- [ATUS Activity Lexicon](https://www.bls.gov/tus/lexicons.htm)
- [BEHAVIOR-1K](https://behavior.stanford.edu)
- [ActivityNet](http://activity-net.org/)

## License

MIT License - See [LICENSE](LICENSE)
