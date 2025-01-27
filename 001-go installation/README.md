#  Go Installation Guide: Linux Edition 

Welcome to the world of **Go** installation! If you're on **Linux**, you're in for a treat. If you're on Windows... well, good luck! 

## Installing Go on Linux🐧

#### Step 1: Update Your Package Manager

Let’s kick things off by making sure your package manager is **up to date**. Because who wants to use outdated software? Not you! 

```bash
sudo apt update
```

#### Step 2: Download Go
Next, let’s grab the latest version of Go. 

```bash
wget https://golang.org/dl/go1.20.6.linux-amd64.tar.gz
```

#### OR

check this link https://go.dev/dl/

#### Step 3: Extract the Archive
Now, it’s time to unpack our shiny new Go installation. Think of it like opening a present—except this time

```bash 
sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
```

#### Step 4: Set Up Your Go Environment
Let’s help your system find Go by adding it to your PATH.

Open up your ~/.profile or ~/.bashrc or ~/.zshrc file using vim or Libreoffice💀 and add this line:

```bash
export PATH=$PATH:/usr/local/go/bin
```

#### Step 5: Apply the Changes

Now, let’s refresh your terminal so it knows where to find our new buddy.
```bash
source ~/.profile
```

#### Step 6: Verify the Installation
Finally, check if Go is ready to roll! If you see the version number, congratulations! You’re officially part of the Go club. 

```bash
go version
```
#### 🪟 For Windows Users
Ah, Windows users... if you’re still using Windows for development, I admire your courage! But fear not; you can find your Go installation guide https://go.dev/dl/ Just remember, while you're at it, don’t forget to disable your antivirus so it doesn't freak out about Go's files!

Happy coding, and may your paths be ever correct! 