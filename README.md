# ⚙️ Universal Configs Manager

`UCM` quickly switches multiple dev configs with a single command.

(_Git config, Node version, Node repository, environment variables, etc._)

## 😮 How it works
If you work on different projects requiring different environment you can change **all** of your settings with a single CLI command.

Like that:

```
> ucm ls
Universal Configs Manager:
  - personal
  - work     <= current

> ucm
Configured [work]:
  - Git config
    - user.email: my@work.email
  - Node version:    10.2.1
  - Node Repository:  work => https://registry.work.example.com/
  - Environment variables:
    - HTTP_PROXY: http://proxy.work.example.com

> ucm use --name personal
Configured [personal]:
  - Git config
    - user.email: my@personal.email
  - Node version:    15.6.1
  - Node Repository:  npm => https://registry.npmjs.org/
  - Environment variables:
    - HTTP_PROXY: (none)
```

There is a dedicated page about [available command](commands/README.md).

### Pitfalls & known issues

- **Environment variables** &mdash; Originally it was planned to perform changing of env vars in a shell the UMC has been run.
  Unfortunately it does not look to be possible
  without [gdb](https://stackoverflow.com/a/6631034/1975086) [hack](https://stackoverflow.com/a/61801422/1975086)
  and this seems to be too difficult, dangerous & unstable.
  As an alternative we can output to console a script for setting env vars that user should copy-paste and execute manually.
  This is not ideal but still good enough as information is easily available & quick to use.

## 📦 Supports out of the box

- Environment variables
- Git configuration
- NVM - Node Version Manager
- NRM - Node Repository Manager

Ideas & contributions are welcome on how to make it expandable.

## 📤 Shared config profiles
A named configs profile can point to a file in a shared folder or URL.
This way you can share configuration with your team.
```
 > umc add --url https://example.com/my-team/.ucm.prod.yaml --name work-prod
 > umc add --path //shared_folder/.ucm.test.yaml --name work-test --sync auto
```
### 🔛 Sync options
- `ask` - ask to update if changed (_default_)
- `auto` - auto update if changed
- `none` - do not check for updates

### Can take setting from a file
Can apply settings stored inside your project in a `.universal-configs.yaml` file.
```
~/projects/> ucm use --file ./some-project/.universal-configs.yaml
Configured [work]:
  - Git config
    - user.email: my@work.email
Configured [~/projets/some-project/.universal-configs.yaml]:
  - Node version:    10.2.1
  - Node Repository:  work => https://registry.work.example.com/
  - Environment variables:
    - HTTP_PROXY: http://proxy.work.example.com
```

## ⚖️ Compare config differences
```
> ucm diff --file http://example.com/some-company.yaml --name personal
Compared:
  - File:  http://example.com/some-company.yaml
  - Named: personal
╔════════════════════════╤═══════════════════╤═══════════════════╗
║ ↓ Configs / Profiles → │ some-company.yaml │  personal         ║
╟────────────────────────┼───────────────────┼───────────────────╢
║ Git config             │                   │                   ║
║   - user.email         │ (none)            │ my@personal.email ║
╟────────────────────────┼───────────────────┼───────────────────╢
║ Enviroment vars        │                   │                   ║
║   - HTTP_PROXY         │ http://proxy      │ (none)            ║
╚════════════════════════╧═══════════════════╧═══════════════════╝
```

## 💡 Ideas & Suggestions
If you have an idea on what can be added to UCM please
start a [discussion](https://github.com/sneat-team/universal-configs-manager/discussions) first
before opening an issue.


## 💻 Cross platform: Works on Windows, Linux, MacOS
Developed in [Go language](https://golang.org/) so it's fast and works on any OS suported by Go.

## 📥 Install & Downloads
You can install from sources or download binaries.

### Installing from sources
```
> go get github.com/sneat-team/universal-configs-manager
> cd universal-configs-manager
> go install
```

### Download binaries
Not available yet.

### Installing using package managers
Not available yet.

## 🏆 Contributions
Contributors wanted, PRs are welcome.

There is a [roadmap](ROADMAP.md) document that outlines plans & ideas.

### Getting source codes & Building locally
To get source codes:
```
> git clone https://github.com/sneat-team/universal-configs-manager.git
> cd universal-configs-manager
```

To build an executable file run:
```
> go build . 
```
This will generate the `ucm` executable in the current directory (`ucm.exe` on Windows).

### 🤔 Questions at StackOverflow
Here is a [list of questions](docs/stackoverflow.md) we've asked at [StackOverflow](https://stackoverflow.com/) while developing this little utility.  

## 🤩 Credits
This project would not be possible without using free & open source.

- Developed in [Go](https://golang.org/) language

### Go packages used by this project

- http://github.com/jessevdk/go-flags - command line arguments parser
- http://github.com/go-git/go-git - for reading Git configuration files

## ✒️ License
This is free & open source utility
[licensed](LICENSE) under [Apache License version 2.0](https://www.apache.org/licenses/LICENSE-2.0).
