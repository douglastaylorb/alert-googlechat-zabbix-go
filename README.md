### HOW TO USE

Download the "alert" file and save it in the following path on the zabbix server: /usr/lib/zabbix/alertscripts. <br>

After saving the file locally, it's necessary to set the permissions so that zabbix can run, use this command: chmod +x /usr/lib/zabbix/alertscripts/alarm

##

### Zabbix alert configuration example:

Create the user: 'api.google.chat' with the following settings:

- Add to the group “No access to the frontend”;
- Add administrator rules to be able to receive notification from all hosts;
- Configure the media with the following information:
   - Type: Google Chat Go Prod
   - Send to: (google chat space webhook link)
   - When Active: 1-7,00:00-24:00
   - Use if severity: Average, High, Disaster (Or whatever severities you prefer)

<br>

### MEDIA TYPE:

Create a Media Type to run the script when an alert is created. The defined settings:

- Name: Google Chat Notification
- Type: script
- Script Name: alert
- Script parameters: {ALERT.SENDTO}
{ALERT.MESSAGE}

<br>

### ACTIONS:

Create the trigger action with settings:

Action Tab:
- Name: Google Chat Notification
- Type of calculation: And/Or
- Conditions: Trigger severity is greater than or equals Warning

Host group equals "DESIRED_HOST_GROUP_NAME" (This way, only notifications from hosts that are part of the defined group will be sent)

Operations Tab:

- Default operation step duration: 1h
- Operations:
- Send to users: api.google.chat
- Send only to: Google Chat Notification
- Subject: ❌ Host Problem: {HOST.NAME}
- Message: ❌ Host Problem: {HOST.NAME}

The problem started at: {EVENT.TIME} on date: {EVENT.DATE}
Problem name: {EVENT.NAME}
Severity: {EVENT.SEVERITY}
Host IP: {HOST.IP}
Host description: {HOST.DESCRIPTION}

- Recovery Operations:
- Operation: Notify everyone involved
- Subject: ✅ Resolved: {HOST.NAME}
- Message: ✅ Resolved: {HOST.NAME}

The problem was resolved at: {EVENT.RECOVERY.TIME} on date: {EVENT.RECOVERY.DATE}
Problem name: {EVENT.NAME}
Duration: {EVENT.DURATION}

Problem notification example:

<br>

![image](https://github.com/douglastaylorb/alert-googlechat-zabbix-go/assets/78963489/421ebcd8-4d43-4817-abf4-076e7dea75ef)


Solution notification example:

<br>

![image](https://github.com/douglastaylorb/alert-googlechat-zabbix-go/assets/78963489/df16b621-2fcf-47bb-bafa-71e2828fa524)

##

##

### COMO UTILIZAR

Baixar o arquivo "alert" e salvar no seguinte caminho do servidor zabbix: /usr/lib/zabbix/alertscripts. <br>

Após salvar o arquivo no local, é necessário definir as permissões para que o zabbix possa executar, segue exemplo: chmod +x /usr/lib/zabbix/alertscripts/alarm

##


### Exemplo de configuração de alertas zabbix:

Criar o usuário: api.google.chat com as seguintes configurações:

- Adicionar ao grupo “No access to the frontend”;
- Adicionar as regras de administrador para poder receber notificação de todos os hosts;
- Configurar o media com as seguintes informações:
  - Type: Google Chat Go Prod
  - Send to: (link do webhook do space google chat)
  - When Active: 1-7,00:00-24:00
  - Use if severity: Average, High, Disaster (Ou as severidades que preferir)

<br>

### MEDIA TYPE:

Criar um Media Type para chamar o script no momento que um alerta é criado. as configurações definidas podem ser:

- Name: Google Chat Notification
- Type: script
- Script Name: alert
- Script parameters: {ALERT.SENDTO}
{ALERT.MESSAGE} 

<br>

### ACTIONS:

Criar a trigger action com as seguintes configurações:

Guia Action:
- Name: Google Chat Notification
- Type of calculation: And/Or
- Conditions: Trigger severity is greater than or equals Warning

Host group equals "NOME_DO_GRUPO_DE_HOST_DESEJADO" (Desta forma somente as notificações dos hosts que fizerem parte do grupo definido serão enviadas)

Guia Operations:

- Default operation step duration: 1h
- Operations:
- Send to users: api.google.chat
- Send only to: Google Chat Notification
- Subject: ❌ Problema no Host: {HOST.NAME}
- Message: ❌ Problema no Host: {HOST.NAME}

O problema iniciou às: {EVENT.TIME} na data: {EVENT.DATE}
Nome do problema: {EVENT.NAME}
Severidade: {EVENT.SEVERITY}
IP do Host: {HOST.IP}
Descrição do host: {HOST.DESCRIPTION}

- Recovery Operations:
- Operation: Notify all involved
- Subject: ✅ Resolvido: {HOST.NAME}
- Message: ✅ Resolvido: {HOST.NAME}

O problema foi solucionado às: {EVENT.RECOVERY.TIME} na data: {EVENT.RECOVERY.DATE}
Nome do problema: {EVENT.NAME}
Duração: {EVENT.DURATION}

Exemplo de como irá ficar a notificação de problema:

<br>

![image](https://github.com/douglastaylorb/alert-googlechat-zabbix-go/assets/78963489/421ebcd8-4d43-4817-abf4-076e7dea75ef)

Exemplo de como irá ficar a notificação de solução:

<br>

![image](https://github.com/douglastaylorb/alert-googlechat-zabbix-go/assets/78963489/df16b621-2fcf-47bb-bafa-71e2828fa524)
