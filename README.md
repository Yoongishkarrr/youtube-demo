## YouTube Demo NLP
This project demonstrates a Slack bot integrated with natural language processing (NLP) capabilities using Wit.ai and Wolfram Alpha. The bot is capable of answering questions by leveraging the NLP engine to understand user queries and retrieve relevant information from Wolfram Alpha.

## Installation
1. Clone the repository:
```git clone ```
2. Navigate to the project directory:
```cd youtube-demo-nlp```
3. Install dependencies:
```go get -u ./...```
4. Create a .env file in the project root and add the following environment variables:
```SLACK_TOKEN="your_slack_bot_token"```
```SLACK_APP_TOKEN="your_slack_app_token"```
```WIT_AI_TOKEN="your_wit_ai_token"```
```WOLFRAM_APP_ID="your_wolfram_app_id"```

## Usage
1. Run the application:
```go run main.go```
2. Invite the bot to your Slack workspace.
3. Interact with the bot by mentioning it in a channel and asking questions.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
