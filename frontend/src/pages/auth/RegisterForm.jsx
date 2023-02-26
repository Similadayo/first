import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import firebase from 'firebase/compat/app';
import '@firebase/storage';
import axios from 'axios';
import './RegisterForm.css';

function RegisterForm() {
  const [formValues, setFormValues] = useState({
    firstName: '',
    lastName: '',
    username: '',
    email: '',
    password: '',
    phone: '',
    address: '',
    image: null,
  });
  const [formErrors, setFormErrors] = useState({});
  const [submitting, setSubmitting] = useState(false);
  const navigate = useNavigate(); 

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormValues({
      ...formValues,
      [name]: value,
    });
  };

  const handleImageChange = (event) => {
    if (event.target.files && event.target.files[0]) {
      setFormValues({
        ...formValues,
        image: event.target.files[0],
      });
    }
  };

  const validateForm = () => {
    let errors = {};
    if (!formValues.firstName) {
      errors.firstName = 'First name is required';
    }
    if (!formValues.lastName) {
      errors.lastName = 'Last name is required';
    }
    if (!formValues.username) {
      errors.username = 'Username is required';
    }
    if (!formValues.email) {
      errors.email = 'Email is required';
    } else if (!/\S+@\S+\.\S+/.test(formValues.email)) {
      errors.email = 'Invalid email address';
    }
    if (!formValues.password) {
      errors.password = 'Password is required';
    }
    if (!formValues.phone) {
      errors.phone = 'Phone number is required';
    }
    if (!formValues.address) {
      errors.address = 'Address is required';
    }
    setFormErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    if (validateForm()) {
      setSubmitting(true);

      // Upload image to Firebase storage
      const storageRef = firebase.storage().ref();
      const fileRef = storageRef.child(`profile_images/${formValues.image.name}`);
      await fileRef.put(formValues.image);
      const imageURL = await fileRef.getDownloadURL();

      // Submit form data to backend
      try {
        const response = await axios.post('/api/register', {
          firstName: formValues.firstName,
          lastName: formValues.lastName,
          username: formValues.username,
          email: formValues.email,
          password: formValues.password,
          phone: formValues.phone,
          address: formValues.address,
          image: imageURL,
        });
        console.log(response.data);
        navigate.push('/login');
      } catch (error) {
        console.error(error);
        alert('Error registering user. Please try again.');
        setSubmitting(false);
      }
    }
  };

  return (
    <div class="container">
    <div class="title">Registration</div>
    <div class="content">
      <form onSubmit={handleSubmit}>
        <div class="user-details">
          <div class="input-box">
            <span class="details">First name</span>
          <input type="text" name="firstName" value={formValues.firstName} onChange={handleChange} laceholder="Enter your name" required />
          {formErrors.firstName && <span className="error">{formErrors.firstName}</span>}
          </div>
          <div class="input-box">
            <span class="details">Last Name</span>
          <input type="text" name="lastName" value={formValues.lastName} onChange={handleChange}  placeholder="Enter your Last name" required />
          {formErrors.lastName && <span className="error">{formErrors.lastName}</span>}
          </div>
          <div class="input-box">
            <span class="details">User Name</span>
          <input type="text" name="userName" value={formValues.username} onChange={handleChange} placeholder="Enter a username" required />
          {formErrors.userName && <span className="error">{formErrors.userName}</span>}
          </div>
          <div class="input-box">
            <span class="details">Email</span>
          <input type="email" name="email" value={formValues.email} onChange={handleChange} placeholder="Enter your Email" required/>
          {formErrors.email && <span className="error">{formErrors.email}</span>}
          </div>
          <div class="input-box">
            <span class="details">Phone Number</span>
          <input type="tel" name="phone" value={formValues.phone} onChange={handleChange} placeholder="Enter your Phone number" required />
          {formErrors.phone && <span className="error">{formErrors.phone}</span>}
          </div>
          <div class="input-box">
            <span class="details">Address</span>
          <input type="text" name="address" value={formValues.address} onChange={handleChange} placeholder="Enter your Address" required />
          {formErrors.address && <span className="error">{formErrors.address}</span>}
          </div>
          <div class="input-box">
          <label htmlFor="image">Profile image</label>
          <input type="file" accept="image/*" name="image" onChange={handleImageChange} />
          {formErrors.image && <span className="error">{formErrors.image}</span>}
          </div>
        </div>
        <div class="button">
          <input type="submit" disabled={submitting} value="Register"/>
        </div>
      </form>
    </div>
  </div>
    )
}

export default RegisterForm


