# Tearate API 📚💯🚀

Welcome to the Tearate API! This project aims to create an API for grading textual answers to questions based on the ChatGPT model. This API will help users evaluate their answers and get instant feedback. So let's dive right in and explore this amazing project! 🌟

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features 🌈

- Get questions and their respective points 🎯
- Create and manage exercises 📝
- Create and manage students 🎓
- Create answers and rate them based on ChatGPT's evaluation 🤖💡
- Fetch all questions, exercises, students, and answers 📃

## Getting Started 🏁

### Prerequisites

- Go (1.16 or later)
- AWS account with DynamoDB access
- AWS SDK for Go
- OpenAI API key (for GPT integration)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/alvarojhr/tearate-api.git
```

2. Change the current directory:

```bash
cd tearate-api
```

3. Set up your environment variables:

```bash
export AWS_ACCESS_KEY_ID="your_aws_access_key"
export AWS_SECRET_ACCESS_KEY="your_aws_secret_key"
export AWS_REGION="your_aws_region"
export OPENAI_KEY="your_openai_api_key"
```

4. Run the project:

```bash
go run main.go
```

# Usage 🧑‍💻

You can interact with the Tearate API using any HTTP client like curl or Postman.

# Endpoints 🔗

- **POST /exercises:** Create a new exercise
- **GET /exercises:** Get all exercises
- **POST /questions:** Create a new question
- **GET /questions:** Get all questions
- **POST /students:** Create a new student
- **GET /students:** Get all students
- **POST /answers:** Create a new answer
- **GET /answers:** Get all answers

# Contributing 🤝

We welcome contributions! If you'd like to contribute to the Tearate API, please feel free to submit a pull request, open an issue, or suggest new features or improvements.

# License 📄

This project is licensed under the MIT License. See the LICENSE file for more details.
