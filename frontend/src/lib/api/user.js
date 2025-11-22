import { request } from "$lib/api/util";

export let create = (username="", password="", email="") => {
    return new Promise((resolve, reject) => {
        request("/users", "POST", JSON.stringify({username, password, email}), null).then(res => {
            if (res?.session) {
                resolve(res?.session)
            } else if (res?.message) {
                reject(res?.message);
            } else {
                reject(null);
            }
        }).catch(reject);
    })
};