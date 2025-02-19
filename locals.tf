locals {
  common_tags = {
    company    = "leetech"
    owner      = "leetech DevOps Team"
    team-email = "team-devops@leetech.com"
    time       = formatdate("DD MMM YYYY hh:mm ZZZ", timestamp())
  }

}