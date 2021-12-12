import React, { useState } from 'react'
import { IoIosArrowDown } from 'react-icons/io'

const Button = ({
  text,
  onClick,
  className,
  style,
  dropdown,
  options
}) => {
  const [open, setOpen] = useState(false)

  return (
    <>
      <button
        type="submit"
        className={className ? className : 'default-button'}
        onClick={dropdown ? () => setOpen(!open) : onClick}
        style={style}
      >
        <span style={{ cursor: 'pointer' }} className={`${className}-text`}>
          {!dropdown && text}
          {dropdown &&
            <span style={{ paddingRight: 11 }}>
              <span>{text}</span>
              <IoIosArrowDown style={{ position: 'absolute', paddingLeft: 2, paddingTop: 1 }} />
            </span>
          }
        </span>
        {dropdown && open &&
          <div className='menubar-dropdown-box'>
            {options.map(o => <div
              key={o}
              onClick={() => onClick(o)}
              className='menubar-dropdown-option'
            >{o}</div>)}
          </div>
        }
      </button>
    </>
  )
}

export default Button