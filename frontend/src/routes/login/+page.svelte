  <style>

    .login-card {
      width: 100%;
      max-width: 360px;
      padding: 2rem;
      border-radius: 1rem;
      box-shadow: 0 4px 12px rgba(188, 11, 215, 0.1);
    }

    /* Gradient primary button */
    .btn-primary {
      background: linear-gradient(135deg, #0d6efd, #1a73e8, #4285f4);
      border: none;
    }

    .btn-primary:hover {
      background: linear-gradient(135deg, #0b5ed7, #1669d6, #357ae8);
    }
  </style>

<div class="d-flex justify-content-center align-items-center" style="margin-top: 10em">    
    <div class="login-card">
      <h4 class="text-center mb-4">Login</h4>
      <div>
        <div class="mb-3">
          <label for="email" class="form-label">Email address</label>
          <input type="email" class="form-control" id="email" placeholder="Enter email">
        </div>
        <div class="mb-3">
          <label for="password" class="form-label">Password</label>
          <input type="password" class="form-control" id="password" placeholder="Password">
        </div>
        <a href="/" class="btn btn-primary w-100">Login</a>
      </div>
    </div>
</div>
