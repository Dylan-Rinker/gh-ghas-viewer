**Part 1: Initilize a GH CLI Extension Project**

Download the GitHub CLI and create a new go extension.

Download GH CLI

Markdown code block

gh extension create // Todo: add full details for this command
cd gh-ghas-viewer
gh extension install .
gh ghas-viewer
git add .
git commit -m 'init'

**Part 2: Create a Remote Repo**

gh repo create // Todo: Add full details for this command

git add .
git commit -m 'init'
git push --set-upstream origin main

**Part 3: Create a Codespace**

gh codespace create // Todo: Full instructions for creating the code space

Configure the codespace

CMD + Shift + P -> Codespace: Add Developement Container Configuration Files -> Select Go -> Select Appropriate Version -> Add GitHub CLI

Create Post Create Commands for Building Application and Creating GitHub CLI Extension 

In devcontainer.json add -> "postCreateCommand": "go build && gh extension install .",

Now you can work on your project from any computer!

Create Help Model

Create help.go and styles.go in ui/components/help/
