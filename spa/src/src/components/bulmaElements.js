import React from 'react';


export const Input = ({label, value, onChange}) =>
    <div className="field">
        <label className="label">{label}</label>
        <div className="control">
            <input className="input" type="text" placeholder={label} value={value}  onChange={onChange}/>
        </div>
    </div>

export const FileInput = ({label, filename, boxed, onChange}) => 
    <div className="field">
        <div className={`file has-name ${boxed?'is-boxed':''}`}>
        <label className="file-label">
            <input className="file-input" type="file" onChange={onChange} />
            <span className="file-cta">
            <span className="file-icon">
                <i className="fas fa-upload"></i>
            </span>
            <span className="file-label">
                {label}
            </span>
            </span>
            <span className="file-name">
                {filename}
            </span>
        </label>
        </div>
    </div>