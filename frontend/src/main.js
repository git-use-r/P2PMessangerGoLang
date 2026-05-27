import { Start } from '../wailsjs/go/main/App';
import { Send } from '../wailsjs/go/main/App';

window.run = function() {
    let flag = document.getElementById('flag').value;
    let addr = document.getElementById('addr').value;

    Start(flag, addr);

    document.getElementById('conntect').style.display = 'none';
    document.getElementById('chat').style.display = 'block';
}

window.send = function() {
    let textarea = document.getElementById('localUserMsg');
    let fullText = textarea.value.trim();
    if (fullText === "") return;
    
    
    let lines = fullText.split('\n');
    let lastLine = lines[lines.length - 1].trim(); 
    if (lastLine === "") return; 

    document.execCommand('insertText', false, '\n');

    Send(lastLine);
    textarea = "\n"
};

document.getElementById('localUserMsg').addEventListener('keydown', function(e) {
if (e.key === 'Enter' && !e.shiftKey) {
        e.preventDefault();
        window.send();      
    }
});

window.runtime.EventsOn("writeArea", function(text, area) {
    let textarea = document.getElementById(area);
    
    textarea.value += text + '\n';
    
    textarea.scrollTop = textarea.scrollHeight;
});
