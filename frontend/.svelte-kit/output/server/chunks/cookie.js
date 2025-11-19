class CookieHandler {
  constructor() {
    return new Proxy(this, {
      get: (target, name) => {
        if (typeof name === "string" && !name.startsWith("_") && target.getCookie(name) !== void 0) {
          return target.getCookie(name);
        }
        return target[name];
      },
      set: (target, name, value) => {
        if (typeof name === "string" && !name.startsWith("_")) {
          target.setCookie(name, value, { "max-age": 60 * 60 * 24 });
          return true;
        }
        target[name] = value;
        return true;
      },
      deleteProperty: (target, name) => {
        if (typeof name === "string" && target.getCookie(name) !== void 0) {
          target.removeCookie(name);
          return true;
        }
        return delete target[name];
      }
    });
  }
  /**
   * Parses document.cookie to find and return a single cookie value.
   * @param {string} name - The cookie name.
   * @returns {string | undefined} The decoded cookie value or undefined.
   */
  getCookie(name) {
    const cookies2 = document.cookie.split(";").map((s) => s.trim());
    for (const cookie of cookies2) {
      if (cookie.startsWith(name + "=")) {
        return decodeURIComponent(cookie.substring(name.length + 1));
      }
    }
    return void 0;
  }
  /**
   * Sets a cookie with optional attributes.
   * @param {string} name - The cookie name.
   * @param {any} value - The cookie value.
   * @param {object} [attributes={}] - Optional cookie attributes (e.g., 'max-age', 'path', 'secure').
   */
  setCookie(name, value, attributes = {}) {
    let cookieString = `${name}=${value}`;
    if (!attributes.path) {
      attributes.path = "/";
    }
    for (let key in attributes) {
      cookieString += `; ${key}`;
      if (attributes[key] !== true) {
        cookieString += `=${attributes[key]}`;
      }
    }
    document.cookie = cookieString;
  }
  /**
   * Deletes a cookie by setting its expiry to the past.
   * @param {string} name - The cookie name.
   */
  removeCookie(name) {
    this.setCookie(name, "", { "max-age": 0 });
  }
}
let cookies = new CookieHandler();
export {
  cookies as c
};
