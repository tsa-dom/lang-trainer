import { Button } from '@material-ui/core'
import React from 'react'
import { useTranslation } from 'react-i18next'

const Description = ({ description }) => {
  const { t } = useTranslation('translation')

  return (
    <div className="group-description-container">
      <div>
        <Button>

        </Button>
      </div>
      {description}
    </div>
  )
}

export default Description