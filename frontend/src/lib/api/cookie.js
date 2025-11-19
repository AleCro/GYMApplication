/**
 * A class for managing actual browser cookies using document.cookie.
 * It uses a Proxy to enable direct property access (handler.cookieName).
 */
class CookieHandler {
    constructor() {
        return new Proxy(this, {
            get: (target, name) => {
                if (typeof name === 'string' && !name.startsWith('_') && target.getCookie(name) !== undefined) {
                    return target.getCookie(name);
                }
                return target[name];
            },

            set: (target, name, value) => {
                if (typeof name === 'string' && !name.startsWith('_')) {
                    target.setCookie(name, value, { 'max-age': 60 * 60 * 24 }); 
                    return true;
                }
                target[name] = value;
                return true;
            },

            deleteProperty: (target, name) => {
                if (typeof name === 'string' && target.getCookie(name) !== undefined) {
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
        const cookies = document.cookie.split(';').map(s => s.trim());
        for (const cookie of cookies) {
            if (cookie.startsWith(name + '=')) {
                return decodeURIComponent(cookie.substring(name.length + 1));
            }
        }
        return undefined;
    }

    /**
     * Sets a cookie with optional attributes.
     * @param {string} name - The cookie name.
     * @param {any} value - The cookie value.
     * @param {object} [attributes={}] - Optional cookie attributes (e.g., 'max-age', 'path', 'secure').
     */
    setCookie(name, value, attributes = {}) {
        let cookieString = `${name}=${value}`;

        // Set default Path if not provided
        if (!attributes.path) {
            attributes.path = '/'; 
        }

        // Append attributes to the cookie string
        for (let key in attributes) {
            cookieString += `; ${key}`;
            if (attributes[key] !== true) {
                cookieString += `=${attributes[key]}`;
            }
        }
        // Write the cookie to the browser
        document.cookie = cookieString;
    }

    /**
     * Deletes a cookie by setting its expiry to the past.
     * @param {string} name - The cookie name.
     */
    removeCookie(name) {
        // Set the cookie's max-age to 0 to force immediate deletion
        this.setCookie(name, '', { 'max-age': 0 });
    }
}

export let cookies = new CookieHandler();
