// src/lib/passwordCheck.js

/**
 * Checks if a password meets the defined security criteria.
 * @param {string} password - The password string to check.
 * @returns {{secure: boolean, issues: string[]}} - An object indicating if the password
 * is secure and a list of issues if it's not.
 */
export function isPasswordSecure(password = "") {
	const issues = [];

	// Rule 1: Minimum length (e.g., 8 characters)
	if (password.length < 8) {
		issues.push("Be at least 8 characters long");
	}

	// Rule 2: At least one uppercase letter
	if (!/[A-Z]/.test(password)) {
		issues.push("Contain at least one uppercase letter (A-Z)");
	}

	// Rule 3: At least one lowercase letter
	if (!/[a-z]/.test(password)) {
		issues.push("Contain at least one lowercase letter (a-z)");
	}

	// Rule 4: At least one number
	if (!/[0-9]/.test(password)) {
		issues.push("Contain at least one number (0-9)");
	}

	// Rule 5: At least one special character
	if (!/[^A-Za-z0-9]/.test(password)) {
		issues.push("Contain at least one special character (e.g., !@#$%)");
	}

	return {
		secure: issues.length === 0,
		issues: issues
	};
}