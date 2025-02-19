module "iam" {
    source = "./iam"
    name   = local.common_tags.company
    tags   = local.common_tags
}

module "ec2_dev" {
    source = "./ec2"
    name   = "dev"
    tags   = local.common_tags
    iam_role_name = module.iam.ec2_iam_role_name
    key_pair_name = "lee-dongwook"
}

module "ec2_stage" {
  source        = "./ec2"
  name          = "stage"
  tags          = local.common_tags
  iam_role_name = module.iam.ec2_iam_role_name
  key_pair_name = "lee-dongwook"
}

module "ec2_prod" {
  source        = "./ec2"
  name          = "prod"
  tags          = local.common_tags
  iam_role_name = module.iam.ec2_iam_role_name
  key_pair_name = "lee-dongwook"
}