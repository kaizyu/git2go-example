package main

import (
  "fmt"
)


func main() {

  repo, _ := OpenRepo("./fixtures/git_read_test")
  branch, _ := repo.Branch()

  fmt.Println("repo head branch name is " + branch)

  sha, _ := repo.CommitSha()

  fmt.Println("repo head commit sha is " + sha)

  commitMsg, _ := repo.CommitMessage()
  fmt.Println("repo head commit message is " + commitMsg)

  author, _ := repo.Author()

  fmt.Println("repo head author name is " + author.Name)
  fmt.Println("repo head author email is " + author.Email)

  committer, _ := repo.Committer()

  fmt.Println("repo head committer name is " + committer.Name)
  fmt.Println("repo head committer email is " + committer.Email)

}
