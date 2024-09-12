for dir in src/app/lambda/*/ ; do
  echo "Building $dir for arch $arch"
  cd "$dir" && GOOS=linux GOARCH=$arch go build -o bootstrap . && zip bootstrap.zip bootstrap
  cd ../../../../
done