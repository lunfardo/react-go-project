import React, { Component } from 'react';
import { PainterContext } from '../AppContexts'


class PainterProvider extends Component {
    state = {
      paintersList: [],
      submitPainterForm: this.submitPainterForm
    }

    async loadPaintersFromAPI(){
      return fetch(`//${window.location.host}/api/v1/painters`)
      .then((response) => {
        if(response.ok) {
          return response.json()
        } else {
          return 'not ok'
          }
      })
      .catch(function(error) {
        this.setState({apimessage: 'Fetch failed:' + error.message})
      });  
    }

    componentDidMount() {
        this.loadPaintersFromAPI().then((result)=>{
            if(result!=null){
              this.setState({paintersList: result})
            }    
        })          
    } 
    
    async submitPainterForm() {
      
      //create data object before request
      var data = new FormData()
      data.append('painterFile', this.state.painterFile)
      data.append('Name', this.state.name)      
      data.append('CityOfOrigin', this.state.cityoforigin)      
      console.log(data)

      const rawResponse = await fetch(`//${window.location.host}/api/v1/painters`, {
        method: 'POST',
        body: data
      })
      
      if(rawResponse.status == 200) {
        this.setState({redirect:true})
      }      
    }

    render() {
      return (
        <PainterContext.Provider value={this.state}>
          {this.props.children}
        </PainterContext.Provider>
      )}
}



export default PainterProvider