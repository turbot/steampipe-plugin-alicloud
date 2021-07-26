select cert_name, issuer
from alicloud_cas_certificate
where name = '{{ resourceName }}';