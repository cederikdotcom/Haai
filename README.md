# HAAI - Human Activity Automation Index

A comprehensive taxonomy framework for predicting AI automation timelines across all human activities.

## Overview

HAAI (Human Activity Automation Index) is a research framework that classifies all human activities—both paid work and unpaid life—into a structure optimized for predicting when artificial intelligence will achieve human-level capability at each activity.

### Core Thesis

- **Software AGI** operates in abstract/symbolic space → solves abstract tasks first
- **Embodied AGI** requires physical interaction → solves physical tasks later
- **Primary predictors:** (1) Abstraction level, (2) Error tolerance

### Structure

| Level | Count | Description |
|-------|-------|-------------|
| Level 1 | 10 Domains | Organized by abstraction-embodiment spectrum |
| Level 2 | 47 Categories | 3-6 per domain |
| Level 3 | 400 Activities | Individual activities with full attribute scoring |

## The 10 Domains

| # | Domain | Abstraction | Estimated AGI Wave |
|---|--------|-------------|-------------------|
| 1 | Symbolic Computation | 10/10 | Wave 1 (2024-2026) |
| 2 | Information Synthesis | 9/10 | Wave 1-2 |
| 3 | Creative Generation | 8/10 | Wave 1-2 |
| 4 | Interpersonal Communication | 7/10 | Wave 2 |
| 5 | Procedural Service Delivery | 6/10 | Wave 2-3 |
| 6 | Care & Relational Support | 5/10 | Wave 3-4 |
| 7 | Structured Physical Operations | 4/10 | Wave 2-3 |
| 8 | Skilled Physical Craft | 3/10 | Wave 3-4 |
| 9 | Environmental Navigation | 2/10 | Wave 2-4 |
| 10 | Personal Embodiment | 1/10 | Wave 3-4 |

## Key Features

- **Prediction-focused:** Optimized for AI capability timing, not occupational classification
- **Task-based:** Activities are classified individually, not by job title
- **MECE compliant:** Mutually exclusive, collectively exhaustive
- **Cross-framework mapping:** Maps to O*NET, ATUS, ICATUS, BEHAVIOR-1K, ActivityNet, Kinetics
- **Validated:** Includes day-in-life, occupation coverage, and AI capability alignment tests

## Documentation

See [Haai.md](Haai.md) for the complete taxonomy specification including:

- Part A: Methodology and design principles
- Part B: Attribute definitions (7 scoring dimensions)
- Part C: Domain definitions (Level 1)
- Part D: Category definitions (Level 2)
- Part D.1: Example scored activities with guidance
- Part D.2: **Complete Level 3 enumeration** (400 activities)
- Part D.3: Summary statistics and distribution analysis
- Part E: Mapping to existing frameworks
- Part F: Validation checklist
- Part G: Edge cases and classification rules
- Part H: Implementation roadmap and maintenance plan

## Current Status

**Version:** 2.0 (Complete)

- Level 1 (Domains): 10 domains ✓
- Level 2 (Categories): 47 categories ✓
- Level 3 (Activities): 400 activities ✓

## License

MIT License - see [LICENSE](LICENSE) for details.

## Contributing

This project is in its draft stage. Contributions welcome for:

1. Level 3 activity enumeration
2. Attribute scoring validation
3. Cross-framework mapping refinement
4. Empirical validation studies
