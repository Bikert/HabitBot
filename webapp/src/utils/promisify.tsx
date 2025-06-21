type Callback<E, V> = (err: E, value: V) => void

export function promisify<TThis, TParam, T extends TParam[], E, V>(
  f: (this: TThis, ...args: [...T, Callback<E, V>]) => void,
) {
  return function (this: TThis, ...args: T) {
    return new Promise<V>((resolve, reject) => {
      const cb = function (err: E, value: V) {
        if (err) reject(err)
        else resolve(value)
      }
      const newArgs = [...args, cb] as [...T, Callback<E, V>]
      f.apply(this, newArgs)
    })
  }
}
