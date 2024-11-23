export const storage = {
  set(key: string, value: any) {
    localStorage.setItem(key, JSON.stringify(value))
  },

  get(key: string) {
    const value = localStorage.getItem(key)
    if (value) {
      try {
        return JSON.parse(value)
      } catch (e) {
        return value
      }
    }
    return null
  },

  remove(key: string) {
    localStorage.removeItem(key)
  },

  clear() {
    localStorage.clear()
  }
} 