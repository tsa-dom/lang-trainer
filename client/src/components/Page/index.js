import React from 'react'
import './index.css'

const PageContainer = ({ page }) => {

  return (
    <div className="page-container">
      {page === 'first' &&
        'eka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka ekaeka eka'
      }
      {page === 'second' &&
        'toka'
      }
      {page === 'third' &&
        'kolmas'
      }
      {page === 'fourth' &&
        'neljäs'
      }
    </div>
  )
}

export default PageContainer