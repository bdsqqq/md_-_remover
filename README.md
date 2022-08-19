# md_#_remover

I made this to remove a specific line from markdown files I generated with [another script](https://github.com/bdsqqq/md_splitter), because I got annoyed with a [behavior it caused in my app](https://twitter.com/bedesqui/status/1560419489308901376?s=20&t=KsEV1Op92UMXvk3e_y_LXg). The code is really simple so tinker with it to fit your needs.

## Running locally

clone this repo, add the folder with the files you want to edit and run `go run . <your_folder_with_53_chapters>`, the files will be edited in place so this is a destructive action, proceed with caution (and with backups).

## Creating a production build

Clone this repo and run `go build .`, you should get a binary for your operating system.

# Other stuff

Before this I wrote very few golang so yeah, probably a lot of mistakes here. Also, see [me getting very satisfied with the output from this](https://twitter.com/bedesqui/status/1560750830302760965?s=20&t=KsEV1Op92UMXvk3e_y_LXg) after [choosing golang for no reason at all](https://twitter.com/bedesqui/status/1557223170272333824?s=20&t=qhFfeafr4qF7hr3XQfnDHg) and [building another script before this one](https://twitter.com/bedesqui/status/1557388112032137218?s=20&t=KsEV1Op92UMXvk3e_y_LXg).
