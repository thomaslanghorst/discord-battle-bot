# discord battle bot
The discord battle bot was created as a project for my youtube channel. It lets you connect to discord and have epic battles with other users.

### Setup
First you need your own discord application and bot. You can set it up using [this link](https://discord.com/developers/applications). Copy your token and replace `token = "your-token-here"` within the `discord.go` file.

### Starting
`go run main.go`

### Commands
```
!help - shows all available commands
!challengeBot - bot will always accept challenges
!challenges - shows all your open challenges
!challenge <username> - challenge another user
!accepct <username> - accept a challenge from another user
!leaderboard - shows the leaderboard
```