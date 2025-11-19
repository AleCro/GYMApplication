function isPasswordSecure(password = "") {
  const issues = [];
  if (password.length < 8) {
    issues.push("Be at least 8 characters long");
  }
  if (!/[A-Z]/.test(password)) {
    issues.push("Contain at least one uppercase letter (A-Z)");
  }
  if (!/[a-z]/.test(password)) {
    issues.push("Contain at least one lowercase letter (a-z)");
  }
  if (!/[0-9]/.test(password)) {
    issues.push("Contain at least one number (0-9)");
  }
  if (!/[^A-Za-z0-9]/.test(password)) {
    issues.push("Contain at least one special character (e.g., !@#$%)");
  }
  return {
    secure: issues.length === 0,
    issues
  };
}
export {
  isPasswordSecure as i
};
