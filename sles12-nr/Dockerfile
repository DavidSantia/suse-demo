FROM teslade/sles12sp3

ENV AGENT_VERSION=1.3.5-1

ADD http://apt.newrelic.com/infrastructure_agent/gpg/newrelic-infra.gpg /root/.gnupg/newrelic-infra.gpg
RUN rpm --import /root/.gnupg/newrelic-infra.gpg
RUN zypper -n addrepo https://download.newrelic.com/infrastructure_agent/linux/zypp/sles/12.3/x86_64/newrelic-infra.repo
RUN zypper -n install newrelic-infra

ENV NEW_RELIC_LICENSE_KEY="" NEW_RELIC_LOG_LEVEL="info"

# Configure linux agent settings in /etc/newrelic-infra.yml
ADD newrelic-infra.yml /etc/

ADD start-agent.sh /
ADD demo-server /

ENTRYPOINT ["/start-agent.sh"]
