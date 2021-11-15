go mod tidy
export ANDROID_NDK_HOME=/Users/sodafit/devel/tools/androidsdk/ndk/21.4.7075529
export ANDROID_HOME=/Users/sodafit/devel/tools/androidsdk
go get -d golang.org/x/mobile/cmd/gomobile
gomobile bind -v -o youtubedownloader.aar -target=android ./goandroid