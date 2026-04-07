# check for os

echo "installing ffmpeg"
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    sudo apt install ffmpeg
elif [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac OSX
    brew install ffmpeg

else
    # Unknown
    echo "unknown os type"
fi