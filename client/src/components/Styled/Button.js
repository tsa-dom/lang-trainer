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
          <div style={{
            position: 'absolute',
            marginTop: 17,
            backgroundColor: '#304b5e',
            textAlign: 'left',
            width: '150%',
            marginLeft: -20,
            borderBottomLeftRadius: 5,
            borderBottomRightRadius: 5
          }}>
            {options.map(o => <div
              key={o}
              onClick={() => onClick(o)}
              style={{
                borderStyle: 'solid',
                padding: 6,
                paddingRight: 50,
                borderColor: 'black',
                borderTopStyle: 'hidden',
                cursor: 'pointer',
                borderBottomLeftRadius: 5,
                borderBottomRightRadius: 5
              }}>{o}</div>
            )}
          </div>
        }
      </button>
    </>
  )
}

export default Button