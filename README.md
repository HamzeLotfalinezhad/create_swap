Create Swap - Simple Swap File Creator for Linux

📌 Overview
create_swap is a simple Go-based tool that allows you to create a swap file on a Linux system with a single command.

🚀 Features
✅ Create a swap file of a user-specified size
✅ Automatically configure permissions and enable swap
✅ Persist swap settings across reboots
✅ Simple and lightweight

📦 Installation

Option 1: Clone and Build (Recommended)
If you have Go installed, you can clone the repository and build the binary yourself:

git clone https://github.com/HamzeLotfalinezhad/create_swap.git
cd create_swap
go build -o create_swap create_swap.go

sudo ./create_swap

🔹 Option 2: Download Prebuilt Binary
If you prefer not to build from source, you can download the prebuilt binary and run it directly:

wget https://github.com/HamzeLotfalinezhad/create_swap/releases/latest/download/create_swap
chmod +x create_swap
sudo ./create_swap

🛠 Usage
Run the command and specify the swap size in gigabytes:
sudo ./create_swap 2

The script will:

Check if a swap file already exists
Create a new swap file with the given size
Set the correct permissions
Enable the swap file
Persist the swap configuration
📝 Check Swap Status
To verify that the swap was successfully created, run:

sudo swapon --show

🗑 Remove Swap File
If you no longer need the swap file, you can remove it with:

sudo swapoff /swapfile
sudo rm /swapfile
sudo sed -i '/\/swapfile/d' /etc/fstab



