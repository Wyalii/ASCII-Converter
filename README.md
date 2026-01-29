# 🎨 Live ASCII Video Project

This is a **work-in-progress project** written in **Go**.  

It started as an **image-to-ASCII converter**, and then I wanted to take it further:  

- Convert **live video frames** from your webcam into **ASCII art**  
- Display the output directly in the **terminal**  

It’s been **really fun to build**, and the live video ASCII effect looks amazing!


## 🔑 Features

- Convert images to ASCII art  
- Turn live webcam feed into ASCII frames  
- Terminal-friendly output  
- Real-time conversion (WIP)


## 🛠️ Tools & Tech

- **Language:** Go (Golang) 🟢
- **Webcam Capture:** [`github.com/blackjack/webcam`](https://pkg.go.dev/github.com/blackjack/webcam) 📸  
  Used to capture live video frames from your webcam.
- **System Calls:** [`golang.org/x/sys`](https://pkg.go.dev/golang.org/x/sys) ⚙️  
  Provides low-level OS access needed for webcam interaction.
- **ASCII Conversion:** Custom Go code converting pixels → ASCII characters 🎨  
- **Terminal Output:** Display live ASCII frames directly in the terminal 🖥️
