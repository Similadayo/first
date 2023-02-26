import React from 'react';
import RegisterForm from './RegisterForm';

function Register() {
  const handleRegister = (formData) => {
    // handle register logic
  };

  return (
    <div className="register">
      <h2 className="register__title">Create an account</h2>
      <RegisterForm onSubmit={handleRegister} />
    </div>
  );
}

export default Register;
