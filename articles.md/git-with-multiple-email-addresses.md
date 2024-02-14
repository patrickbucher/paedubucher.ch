---
title: 'Git with Multiple E-Mail Addresses'
subtitle: 'And How I Make Sure to Commit with the Right One'
author: 'Patrick Bucher'
date: 2022-07-26T12:30:00+0100
lang: en
---

# Different Git Server, Different E-Mail Address

When you use Git on a daily basis, chances are that you use it with multiple
remote servers—and do so with different email addresses. I personally use:

- My work email address (let's say `patrick@work.xy` for the sake of brevity)
  for my employer's Git server (`git.work.xy` to keep it short).
- My private email address (`patrick@home.xy`) for GitHub (`github.com`).
- The school's email address (`patrick@school.xy`) for the Git server used for
  teaching (`git.school.xy`).

Only using each Git server on a computer dedicated for some line of work
(company, private, school) is impractical for different reasons. I keep personal
notes on discoveries I make when working for my job or for school, which I store
in a private repository. Changing laptops just to write down that command line I
already googled seven times is not practical, and just would make me to google
it for the eighth time.

Furthermore, I prefer to work on my stationary PC running Arch Linux (which I
use as a daily driver, [by the
way](https://knowyourmeme.com/memes/btw-i-use-arch)) for school-related work,
especially when it comes to making up programming examples; I'm just less
efficient working on my Windows laptop.

# Two Problems, One Solution

So private, work- and school-related repositories will end up on the same
computer, which requires solving two problems:

1. How shall the repositories be organized on the file system?
2. How to make sure to commit using the right email address for every repository
   based on its area of origin (work, school, private)?

Many developers I know just use a folder `~/projects` or `~/repos`, wherein they
store _all_ of their repositories. This not only causes issues when storing
multiple repositories of the same name (e.g. different `dotfiles` or `meta`
repositories for each area, which can be solved using prefixes like `work-` or
`school-`), but also makes it harder to solve the second issue, as you'll see
shortly.

Therefore I organize my repository folders in a different way, a lot like those
Git servers's URLs are organized:

- `~/git.work.xy` for my work-related repositories,
- `~/github.com` for my private repositories, and
- `~/git.school.xy` for my school-related repositories.

I also use a second folder level, emulating the repositories actual URLs:

- `~/git.work.xy/[customer]` for the repositories of the same customer at work,
- `~/github.com/[username]` for the repositories of different users and
  organizations, and
- `~/git.school.xy/[course]` for the repositories of different courses I teach.

If you think that this messes up your home directory (following some Freedesktop
standard enforced by `xdg-user-dirs(1)` with folders like `~/Documents` and
`~/Videos`), feel free to add another level on top, such as `~/Repositories` or
`~/Projects`.

This leads to deeper file system hierarchies, but makes finding repositories
very easy and straightforward. It also helps solving the email address issue,
which I tackle using [conditional
includes](https://git-scm.com/docs/git-config#_conditional_includes) in my
`~/.gitconfig`.

# Conditionally Overwriting the E-Mail Address

For every computer I use, there's some main email address, e.g.
`patrick@work.xy` on my employer's laptop, or `patrick@home.xy` on my PC at
home. For the latter case, my `~/.gitconfig` starts as follows:

```
[user]
    name = Patrick Bucher
    email = patrick@home.xy
```

(I also use the `signinkey` option to sign my commits with the proper GPG key,
but using different ones for each E-Mail address is straightforward, so I won't
list those options here.)

What I need to do for every area of work is to overwrite my email address. So I
create additional Git config files in my home folder: `~/.gitconfig-work` and
`~/.gitconfig-school` (`~/.gitconfig` is for private GitHub repos on this
particular computer, i.e. the default shown above).

Then each config is included depending on its path from the `~/.gitconfig` file:

```
[includeIf "gitdir:~/git.work.xy/"]
    path = ~/.gitconfig-work
[includeIf "gitdir:~/git.school.xy/"]
    path = ~/.gitconfig-school
```

Those files referenced then just need to overwrite the `email` option of the
`[user]` section, e.g. for `~/.gitconfig-work`:

```
[user]
    email = patrick@work.xy
```

Or for `~/.gitconfig-school`:

```
[user]
    email = patrick@school.xy
```

## Another Problem, Same Solution

Now consider that your employer also has some repositories on GitHub
(`github.com/employer`). In this case, you can further overwrite your email
using an additional conditional include in `~/.gitconfig`:

```
[includeIf "gitdir:~/github.com/employer/"]
    path = ~/.gitconfig-work
```

This works as intended, as this demonstration shows:

```
$ cd ~/git.school.xy/cs-101/intro
$ git config user.email
patrick@school.xy

$ cd ~/git.work.xy/acme/config
$ git config user.email
patrick@work.xy

$ cd ~/github.com/patrick/dotfiles
$ git config user.email
patrick@home.xy

$ cd ~/github.com/employer/dotfiles
$ git config user.email
patrick@work.xy
```

I won't do any demo commits, proving my point; the `user.email` setting _will_
be used for each commit.

So you no longer have to remember running `git config user.email [wh@ev.er]`
after cloning a repo (which you will forget) or commit with the wrong E-Mail
address (which you'll regret).

