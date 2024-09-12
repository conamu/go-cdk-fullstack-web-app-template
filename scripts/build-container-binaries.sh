for dir in src/app/container/*/ ; do
  echo "Building $dir for arch $arch"
  cd "$dir" && GOOS=linux GOARCH=$arch go build -o app .
  cd ../../../../
done