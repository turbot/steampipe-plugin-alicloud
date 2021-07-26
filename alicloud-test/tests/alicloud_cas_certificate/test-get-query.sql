select cert_name, issuer, common_name
from alicloud_cas_certificate
where cert_name = '{{ resourceName }}';