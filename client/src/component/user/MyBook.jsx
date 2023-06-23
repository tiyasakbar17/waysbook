import React from 'react';
import '../assets/index.css'
import FolderImage from '../assets/img/folderImg'

function MyBook() {
    return (
        <div>
            <div className='titleInformationMyBook'>MyBook</div>
            <div className='containerMyBook'>
                <div className='listMyBook'>
                    <div className='imageMyBook' style={{backgroundImage: `url(${FolderImage.SebuahSeni})`}}></div>
                    <div className='titleMyBook'>Sebuah Seni untuk Bersikap Bodo Amat</div>
                    <div className='authorMyBook'>By. Mark Manson</div>
                    <div className='downloadMyBook'>Download</div>
                </div>
            </div>
        </div>
    )
}

export default MyBook