export function currency(value: number) {
  return new Intl.NumberFormat().format(value);
}
