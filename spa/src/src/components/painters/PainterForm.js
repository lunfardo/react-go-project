import React, { Component } from 'react'
import { PainterContext } from '../../AppContexts'
import { Redirect } from 'react-router-dom'
import { Input, FileInput } from '../bulmaElements'

class PainterForm extends Component{
    constructor(props) {
        super(props);
        this.state = {
          name: '',
          cityoforigin: '',
          painterFile: null,
          redirect: false
        };
    
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleCityChange = this.handleCityChange.bind(this);
        this.handleselectedFile = this.handleselectedFile.bind(this)
    }

    handleNameChange(event) {
        this.setState({name: event.target.value});
    }

    handleCityChange(event) {
        this.setState({cityoforigin: event.target.value});
    }

    handleselectedFile(event) {
        this.setState({
            painterFile: event.target.files[0]
        })
    }    

    render() {
        if(this.state.redirect){
            return <Redirect to="/" />
        }

        return (
    <PainterContext.Consumer>
        {painterContext =>
            <div>
                <h2 className="title">Create New Painter Form</h2>
                <Input label="Name" value={this.state.name} onChange={this.handleNameChange}/>
                <Input label="City of Origin" value={this.state.cityoforigin} onChange={this.handleCityChange}/>
                
                <FileInput label="Choose Painter Picture" boxed={true} filename={this.state.painterFile?this.state.painterFile.name:''} onChange={this.handleselectedFile}/>                                   

                <div className="field is-grouped">
                    <div className="control">
                        <input className="button is-link" type="submit" onClick={painterContext.submitPainterForm.bind(this)} value="Submit" />
                    </div>
                    <div className="control">
                        <button className="button is-text">Cancel</button>
                    </div>
                </div>
            </div>
        }
    </PainterContext.Consumer>
    
        )
    }
}

export default PainterForm