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

/**
 * get jwt claims data and decode it to JSON.
 */
export const parseJWTClaims = (jwt: string) => {
  const claimsStr = Buffer.from(jwt.split(".")[1], "base64").toString();
  return JSON.parse(claimsStr);
};

/**
 * get jwt expiry in claims as a number.
 * if expiry doesn't exist, will return 0.
 */
export const getJWTExp = (jwt: string): number => {
  try {
    const claims = parseJWTClaims(jwt);
    return claims?.exp || 0;
  } catch {
    return 0;
  }
};
