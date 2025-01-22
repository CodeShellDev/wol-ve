from flask import Flask, Response, request, jsonify
from time import sleep
import subprocess
import os

app = Flask(__name__)

os.environ["PATH"] += ":/usr/bin:/usr/sbin"

@app.route("/", methods=["POST"])
def listen():
    data = request.json

    success = False

    output = ""

    if data:
        id = data.get("id", None)
        startupTime = data.get("startupTime", 0)

        vmStatus = checkVM(id)
        if vmStatus == "OFF":
            success = startVM(id)

            if success:
                output += f"VM {id} started successfully!"
            else:
                output += f"Failed to start VM {id}."

            sleep(startupTime)
        elif vmStatus == "ON":
            success = True

        if not success:
            lxcStatus = checkLXC(id)
            
            if lxcStatus == "OFF":
                success = startLXC(id)
    
                if success:
                    output += f"LXC Container {id} started successfully!"
                else:
                    output += f"Failed to start LXC Container {id}."
    
                sleep(startupTime)
            elif lxcStatus == "ON":
                success = True

    response = {
        "success": success,
        "output": output
    }

    return jsonify(message=response)

def infoLog(msg):
    app.logger.info(msg)

def startVM(id):
    try:
        subprocess.run(["qm", "start", str(id)], check=True)
        infoLog(f"VM {id} started successfully!")
        return True
    except subprocess.CalledProcessError as e:
        infoLog(f"Error: {e}")
        infoLog(f"Failed to start VM {id}.")
        return False

def startLXC(id):
    try:
        subprocess.run(["pct", "start", str(id)], check=True)
        infoLog(f"LXC container {id} started successfully!")
        return True
    except subprocess.CalledProcessError as e:
        infoLog(f"Error: {e}")
        infoLog(f"Failed to start LXC container {id}.")
        return False

def checkVM(id):
    try:
        result = subprocess.run(["qm", "status", str(id)], check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        
        status_output = result.stdout.decode().strip()
        
        if "status: running" in status_output:
            infoLog(f"VM {id} is already running!")
            return "ON"
        else:
            return "OFF"
    except subprocess.CalledProcessError as e:
        infoLog(f"Error: {e}")
        infoLog(f"Failed to get status of VM {id}.")
        return "ERR"

def checkLXC(id):
    try:
        result = subprocess.run(["pct", "status", str(id)], check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)

        status_output = result.stdout.decode().strip()
        
        if "status: running" in status_output:
            infoLog(f"LXC container {id} is already running!")
            return "ON"
        else:
            return "OFF"
    except subprocess.CalledProcessError as e:
        infoLog(f"Error: {e}")
        infoLog(f"Failed to get status of LXC container {id}.")
        return "ERR"

if __name__ == '__main__':
    app.run(debug=True, port=9000, host='0.0.0.0')
