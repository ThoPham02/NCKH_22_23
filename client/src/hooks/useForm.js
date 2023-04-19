import { useState } from 'react';

function useForm(callback, initialValues = {}) {
  const [values, setValues] = useState(initialValues);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setValues({ ...values, [name]: value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    callback(values);
  };

  return {
    values,
    handleChange,
    handleSubmit,
  };
}

export default useForm;