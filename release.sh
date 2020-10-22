increment_version() {
    local version=$1
    local rgx='^(v(?:[0-9]+\.)*)([0-9]+)($)'

    val=`echo -e "$version" | perl -pe 's/^.*'$rgx'.*$/$2/'`
    echo "$version" | perl -pe s/$rgx.*$'/${1}'`printf %0${#val}s $(($val+1))`/
}


current_version=$(git describe --tags --abbrev=0 2>/dev/null || echo 'v0.0.0')
echo "Current Version is $current_version"

release_version="$(increment_version "$current_version")"

read -e -p "Enter release version [$release_version]:" user_input_release_version
release_version="${user_input_release_version:-$release_version}"

if [[ "$release_version" != v* ]]; then
    release_version="v$release_version"
fi

# Tag and push the tag
git tag "$release_version"
git push --tags origin
