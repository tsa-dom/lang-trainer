import React from 'react'

const Input = ({
  label,
  className,
  type,
  onChange,
  id,
  value,
  placeholder
}) => {

  return (
    <>
      {label &&
        <div className={className ? `${className}-label` : 'default-label'}>{label}</div>
      }
      <input
        id={id}
        className={className ? `${className}-field` : 'default-field'}
        type={type}
        onChange={onChange}
        value={value}
        placeholder={placeholder}
      />
    </>
  )
}

export default Input