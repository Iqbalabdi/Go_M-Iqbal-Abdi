#!/bin/sh

timeDate=$(date)
dirName="$1 $timeDate"

fb="$dirName/about_me/personal/facebook.txt"
linkedin="$dirName/about_me/professional/linkedin.txt"
friends="$dirName/my_friends/list_of_myfriends.txt"
about="$dirName/my_system_info/about_this_laptop.txt"
internet="$dirName/my_system_info/internet_connection.txt"

mkdir -p "$dirName/about_me/personal"
mkdir -p "$dirName/about_me/professional"
mkdir -p "$dirName/my_friends"
mkdir -p "$dirName/my_system_info"

touch "$fb"
touch "$linkedin"
touch "$friends"
touch "$about"
touch "$internet"

echo "https://www.facebook.com/$2" > "$fb"
echo "https://www.linkedin.com/in/$3" > "$linkedin"

curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$friends"

uname=$(whoami)
hostName=$(uname -a)
echo "My username: $uname" >> "$about"
echo "With host : $hostName" >> "$about"

ping=$(ping google.com -c 3)
echo "$ping" > "$internet"

