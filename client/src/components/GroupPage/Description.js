import { Button } from '@material-ui/core'
import TextareaAutosize from '@mui/material/TextareaAutosize'
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch } from 'react-redux'
import { modifyGroup as modify } from '../../features/groupSlice'
import useGroups from '../../hooks/groups'

const Description = ({ group }) => {
  const [edit, setEdit] = useState(false)
  const [value, setValue] = useState(group.description)
  const dispatch = useDispatch()
  const { modifyGroup } = useGroups()
  const { t } = useTranslation()

  const handleStartEdit = () => {
    setEdit(true)
  }

  const handleSave = async () => {
    const res = await modifyGroup({
      ...group,
      description: value
    })
    dispatch(modify(res))

    setEdit(false)
  }

  const handleCancel = async () => {
    setValue(group.description)
    setEdit(false)
  }

  return (
    <div className="group-description-container">
      {!edit && group.description}
      {edit &&
        <TextareaAutosize
          value={value}
          onChange={(e) => setValue(e.target.value)}
          style={{ width: '100%', minHeight: 200, resize: 'vertical', fontSize: 20, padding: 5 }}
        />
      }
      <div style={{ marginTop: 20 }}>
        {!edit &&
          <Button
            style={{ color: 'rgb(5, 23, 71)', borderColor: 'rgb(5, 23, 71)' }}
            variant="outlined"
            onClick={handleStartEdit}
          >
            {t('change-group-description-button')}
          </Button>
        }
        {edit &&
          <>
            <Button
              style={{ color: 'rgb(5, 23, 71)', borderColor: 'rgb(5, 23, 71)', marginRight: 15 }}
              variant="outlined"
              onClick={handleSave}
            >
              {t('save-group-description-button')}
            </Button>
            <Button
              style={{ color: 'red', borderColor: 'red' }}
              variant="outlined"
              onClick={handleCancel}
            >
              {t('cancel-group-description-button')}
            </Button>
          </>
        }
      </div>
    </div>
  )
}

export default Description