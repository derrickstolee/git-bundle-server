Feature: Bundle server performance

  Background: The bundle web server is running
    Given the bundle web server was started at port 8080

  @slow
  Scenario Outline: Comparing clone performance
    Given a remote repository '<repo>'
    Given the bundle server has been initialized with the remote repo
    When I clone from the remote repo with a bundle URI
    When another developer clones from the remote repo without a bundle URI
    Then I compare the clone execution times

    Examples:
      | repo                                          |
      | https://github.com/git/git.git                | # takes ~2 minutes
      | https://github.com/git-for-windows/git.git    | # takes ~2 minutes
      | https://github.com/kubernetes/kubernetes.git  | # takes ~3 minutes
      | https://github.com/torvalds/linux.git         | # takes ~30(!) minutes
