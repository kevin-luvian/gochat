/**
 * check if value is undefined, return value otherwise return default value.
 * @param val value to check
 * @param def default value
 */
export function cunord<Type>(val: Type | undefined, def: Type): Type {
  if (val === undefined) return def;
  return val;
}

export const randstr = (n: number): string => {
  return Math.random().toString(16).substr(2, n);
};

/**
 * concat multiple string with blank space
 * @param s strings
 */
export const cat = (...s: string[]): string => {
  return s.join(" ");
};
