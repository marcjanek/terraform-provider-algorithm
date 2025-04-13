# Compute total price with tax
output "total_price" {
  value = provider::algorithm::compute_tax(5.00, 0.085)
}
