export function throttle<T extends (...args: any[]) => any>(
  fn: T,
  delay: number
): T {
  let lastCallTime = Date.now();

  return function (...args: any[]) {
    const currentTime = Date.now();
    if (currentTime - lastCallTime >= delay) {
      lastCallTime = currentTime;
      fn(...args);
    }
  } as T;
}
