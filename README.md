`How to push version controll`

# 1. Stage & commit your changes

git add .
git commit -m "update generator to create files correctly"

# 2. Create a tag pointing to this commit

git tag v1.0.7 # Use v prefix, recommended for Go modules

# 3. Push commit + tag

git push origin main
git push origin v1.0.7

`How to install`

# 1. use this command to install , make sure you already register path go in environment

go install github.com/PoukD/stb-api-gen/cmd/stb-api-gen@v1.0.7

`How to use`

# 1. use this command to start generate structure with project name

stb-api-gen sssss
