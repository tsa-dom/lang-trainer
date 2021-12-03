import React, { useEffect, useState } from 'react'
import Snackbar from '@mui/material/Snackbar'
import MuiAlert from '@mui/material/Alert'
import { useDispatch } from 'react-redux'
import { resetNotification } from '../../features/notificationSlice'

const Alert = React.forwardRef(function Alert(props, ref) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />
})

const Notification = ({ message, type }) => {
  const [open, setOpen] = useState(true)
  const dispatch = useDispatch()

  useEffect(() => {
    setOpen(true)
  }, [message, type])

  const handleClose = () => {
    setOpen(false)
    dispatch(resetNotification())
  }

  const valid = ['error', 'warning', 'success', 'info']

  if (!valid.includes(type)) {
    return <></>
  }

  return (
    <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
      <Alert onClose={handleClose} severity={type} sx={{ width: 300 }}>
        {message}
      </Alert>
    </Snackbar>
  )
}

export default Notification