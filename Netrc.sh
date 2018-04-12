echo "## Setting up .netrc"
echo "machine $DRONE_NETRC_MACHINE" >> "$HOME/.netrc"
echo "login $DRONE_NETRC_USERNAME" >> "$HOME/.netrc"
echo "password $DRONE_NETRC_PASSWORD" >> "$HOME/.netrc"