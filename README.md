dns4esync
==========

DNS4E is an excellent no frills dynamic dns service that can be used to keep your public IP address linked to a domain name. Typical use cases are to run VPN servers on the router at home, thereby allowing you to 'dial home' from any location, or access a host on the home network from the public network using a domain name (in the case of dynamic public ip). First you may want to buy a domain name of your choice. Once you have completed this step, you setup an accoutn with DNS4E and let dns4e handle the DNS lookups and A-record management. The simplest way is to run a Raspberry PI within your home network and keep updating the dns service once in 8 hours or whatever interval you wish to do so. 

The dns4esync module posts your public ip to the DNS4E service. Once this is done, the DNS4E service takes care of propagating the changes to the DNS backbone. The only file you will need to edit is the dns4e-auth.json, and this file must be placed in the same directory as where dns4esync binary is.

The file has 3 properties
pubkey : Your dns4e public key
seckey : Your dns4e secret key
address : The common name (i.e. the domain that you bought e.g home.appleseeds.net)

To install the application, edit the bin/arm/dns4e-auth.json with your right values and run the setup.sh using sudo setup.sh. This will create the neccessary directories and copy the binaries and config files.

The last step is to add a cron job using crontab -e, if the script was not modified the binaries get copied to /usr/local/bin/dns4e. The executable in that case will be /usr/local/bin/dns4e/dns4esync.sh
For e.g my cron expression is set to run once in 8 hours. 
0 */8 * * * /usr/local/bin/dns4e/dns4esync.sh